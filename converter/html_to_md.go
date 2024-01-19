package converter

import (
	"fmt"
	html2md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/JohannesKaufmann/html-to-markdown/plugin"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"htmltomarkdown/config"
	"htmltomarkdown/models"
	"htmltomarkdown/util"
	"mime"
	"path/filepath"
	"regexp"
	"strings"
)

type GetAbsoluteURL func(selection *goquery.Selection, rawURL string, domain string) string

var rules = []html2md.Rule{{
	Filter: []string{"p", "div"},
	Replacement: func(content string, selection *goquery.Selection, opt *html2md.Options) *string {
		parent := goquery.NodeName(selection.Parent())
		if html2md.IsInlineElement(parent) || parent == "li" {
			return &content
		}

		content = html2md.TrimpLeadingSpaces(content)
		content = "\n" + content + "\n"
		return &content
	},
}}

var replacementRules = []models.RegexReplace{{
	// anchor links are escaped by the converter, this fixes it
	// e.g (\#anchor) -> (#anchor)
	Pattern: "\\(\\\\#([^)]+)\\)",
	With:    "(#$1)",
}}

func HtmlConverter(baseUrl, path string, cfg config.Config) *html2md.Converter {
	conv := html2md.NewConverter(path, true, &html2md.Options{
		GetAbsoluteURL: getAbsoluteURL(baseUrl, cfg.AssetDir),
	})

	conv.Before(remove(cfg.Html.Remove), replace(cfg.Html.Replace))
	conv.After(regexReplace(append(replacementRules, cfg.Markdown.Replace...)))
	conv.Use(plugin.Table(), ArticlePlugin(cfg.Html.HeaderTags, ";;;"))
	conv.AddRules(rules...)

	return conv
}

func getAbsoluteURL(baseUrl, assetDir string) GetAbsoluteURL {
	return func(_ *goquery.Selection, rawURL string, path string) string {
		ext := filepath.Ext(rawURL)
		if len(ext) > 0 && ext != ".html" {
			mimeType := mime.TypeByExtension(ext)
			path := util.PathByType(mimeType, rawURL, assetDir)
			return fmt.Sprintf("{{%%baseurl%%}}/%s", path)
		}

		if strings.Contains(rawURL, "github-issues") {
			log.Warnf("github-issues found: %s", rawURL)
		}

		if util.IsExternalUrl(rawURL) {
			return rawURL
		}

		if strings.HasPrefix(rawURL, "#") {
			return rawURL
		}

		// remove baseUrl from url e.g /docs/article -> /article
		if filepath.IsAbs(rawURL) {
			return strings.TrimPrefix(rawURL, baseUrl)
		}

		filename := filepath.Base(rawURL)
		if strings.HasPrefix(filename, "index.html") { // remove index.html from url
			rawURL = filepath.Join(filepath.Dir(rawURL), strings.TrimPrefix(filename, "index.html"))
		}

		// resolve relative urls against the current path
		absURL := filepath.Join(path, rawURL)
		absURL = strings.Replace(absURL, ".html", "", -1)

		return fmt.Sprintf("{{< ref \"%s\" >}}", absURL)
	}
}

func replace(replacements []models.DocReplacement) html2md.BeforeHook {
	return func(doc *goquery.Selection) {
		for _, replacement := range replacements {
			doc.Find(replacement.Match).Each(func(i int, selection *goquery.Selection) {
				var args []string
				for _, selector := range replacement.Select {
					if selector == "text" {
						args = append(args, selection.Contents().First().Text())
						continue
					} else if strings.HasPrefix(selector, "?") {
						selector = strings.TrimPrefix(selector, "?")
						args = append(args, selection.AttrOr(selector, ""))
						continue
					}

					argSel := selection.Find(selector)
					if argSel == nil {
						continue
					}
					args = append(args, argSel.Contents().First().Text())
				}

				rep := util.ReplaceTemplate(replacement.Replace, args)
				selection.ReplaceWithHtml(fmt.Sprintf("<div>%s</div>", rep))
			})
		}
	}
}

func remove(selectors []string) html2md.BeforeHook {
	return func(doc *goquery.Selection) {
		for _, selector := range selectors {
			selection := doc.Find(selector)
			if selection == nil {
				continue
			}
			selection.Remove()
		}
	}
}

func regexReplace(replacements []models.RegexReplace) html2md.Afterhook {
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
