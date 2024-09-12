package converter

import (
	"fmt"
	"htmltomarkdown/config"
	"htmltomarkdown/models"
	"htmltomarkdown/util"
	"mime"
	"path/filepath"
	"regexp"
	"strings"

	html2md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/JohannesKaufmann/html-to-markdown/plugin"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
)

type GetAbsoluteURL func(selection *goquery.Selection, rawURL string, domain string) string

// Define custom conversion rules
var rules = []html2md.Rule{
	{
		Filter: []string{"p", "div"},
		Replacement: func(content string, selection *goquery.Selection, opt *html2md.Options) *string {
			parent := goquery.NodeName(selection.Parent())
			if html2md.IsInlineElement(parent) || parent == "li" {
				return &content
			}

			content = html2md.TrimpLeadingSpaces(content)
			formattedContent := fmt.Sprintf("\n%s\n", content)
			return &formattedContent
		},
	},
}

// Handle anchor link conversion
var replacementRules = []models.RegexReplace{
	{
		Pattern: "\\(\\\\#([^)]+)\\)", // e.g (\#anchor) -> (#anchor)
		With:    "(#$1)",
	},
}

// Create HTML to Markdown converter
func HtmlConverter(baseUrl, path string, cfg config.Config) *html2md.Converter {
	converter := html2md.NewConverter(path, true, &html2md.Options{
		GetAbsoluteURL: getAbsoluteURL(baseUrl, cfg.AssetDir),
	})

	converter.Before(removeElements(cfg.Html.Remove), applyReplacements(cfg.Html.Replace))
	converter.After(applyRegexReplacements(append(replacementRules, cfg.Markdown.Replace...)))

	// Add plugins and custom rules
	converter.Use(
		plugin.Table(),
		plugin.ConfluenceAttachments(),
		plugin.ConfluenceCodeBlock(),
		ArticlePlugin(cfg.Html.HeaderTags, ";;;"),
		handleNestedLists(),
		ignoreLinksInCodeBlocks(),
	)
	converter.AddRules(rules...)
	return converter
}

// Generate absolute URLs for assets and internal links
func getAbsoluteURL(baseUrl, assetDir string) GetAbsoluteURL {
	return func(_ *goquery.Selection, rawURL string, path string) string {
		ext := filepath.Ext(rawURL)
		if isNonHTMLFile(ext) {
			return handleNonHTMLFile(rawURL, ext, assetDir)
		}

		if strings.Contains(rawURL, "github-issues") {
			log.Warnf("github-issues found: %s", rawURL)
		}

		if util.IsExternalUrl(rawURL) || strings.HasPrefix(rawURL, "#") {
			return rawURL
		}

		return resolveRelativeURL(baseUrl, rawURL, path)
	}
}

// Helper to determine non-HTML files
func isNonHTMLFile(ext string) bool {
	return len(ext) > 0 && ext != ".html"
}

// Handle non-HTML file paths
func handleNonHTMLFile(rawURL, ext, assetDir string) string {
	mimeType := mime.TypeByExtension(ext)
	assetPath := util.PathByType(mimeType, rawURL, assetDir)
	return fmt.Sprintf("{{%%baseurl%%}}%s", assetPath)
}

// Resolve relative URLs
func resolveRelativeURL(baseUrl, rawURL, path string) string {
	// Handle absolute paths
	if filepath.IsAbs(rawURL) {
		return strings.TrimPrefix(rawURL, baseUrl)
	}

	// Remove "index.html" from the URL
	if strings.HasPrefix(filepath.Base(rawURL), "index.html") {
		rawURL = filepath.Join(filepath.Dir(rawURL), "")
	}

	// Resolve and format final URL
	absURL := filepath.Join(path, rawURL)
	absURL = strings.Replace(absURL, ".html", "", -1)
	return fmt.Sprintf("{{< ref \"%s\" >}}", absURL)
}

// Replace content using provided replacement templates
func applyReplacements(replacements []models.DocReplacement) html2md.BeforeHook {
	return func(doc *goquery.Selection) {
		for _, replacement := range replacements {
			doc.Find(replacement.Match).Each(func(i int, selection *goquery.Selection) {
				args := extractReplacementArgs(selection, replacement.Select)
				replacedContent := util.ReplaceTemplate(replacement.Replace, args)
				selection.ReplaceWithHtml(fmt.Sprintf("<div>%s</div>", replacedContent))
			})
		}
	}
}

// Extract arguments for replacement templates
func extractReplacementArgs(selection *goquery.Selection, selectors []string) []string {
	var args []string
	for _, selector := range selectors {
		switch {
		case selector == "text":
			args = append(args, selection.Contents().First().Text())
		case strings.HasPrefix(selector, "?"):
			attr := strings.TrimPrefix(selector, "?")
			args = append(args, selection.AttrOr(attr, ""))
		default:
			if argSel := selection.Find(selector); argSel != nil {
				args = append(args, argSel.Contents().First().Text())
			}
		}
	}
	return args
}

// Remove elements by selector
func removeElements(selectors []string) html2md.BeforeHook {
	return func(doc *goquery.Selection) {
		for _, selector := range selectors {
			doc.Find(selector).Remove()
		}
	}
}

// Apply regex replacements to final markdown output
func applyRegexReplacements(replacements []models.RegexReplace) html2md.Afterhook {
	return func(markdown string) string {
		for _, replacement := range replacements {
			rgx, err := regexp.Compile(replacement.Pattern)
			if err != nil {
				log.Errorf("failed to compile pattern: %s", replacement.Pattern)
				continue
			}
			markdown = rgx.ReplaceAllString(markdown, replacement.With)
		}
		return markdown
	}
}
