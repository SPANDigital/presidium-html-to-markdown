package converter

import (
	"fmt"
	"htmltomarkdown/util"
	"path/filepath"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
)

// ArticlePlugin converts markdown headers to hugo article headers.
// E.g:
// # This is a header
// becomes:
// ---
// title: This is a header
// ---
func ArticlePlugin(filters []string, delim string) md.Plugin {
	return func(c *md.Converter) []md.Rule {
		return []md.Rule{
			{
				Filter: filters,
				Replacement: func(content string, selec *goquery.Selection, opt *md.Options) *string {
					if strings.TrimSpace(content) == "" {
						return nil
					}

					content = strings.Replace(content, "\n", " ", -1)
					content = strings.Replace(content, "\r", " ", -1)
					content = strings.TrimSpace(content)

					insideLink := selec.ParentsFiltered("a").Length() > 0
					if insideLink {
						text := opt.StrongDelimiter + content + opt.StrongDelimiter
						text = md.AddSpaceIfNessesary(selec, text)
						return &text
					}

					title := util.EscapeTitle(selec.Text())
					header := articleHeader(title, delim)
					return &header
				},
			},
		}
	}
}

func articleHeader(title string, delim string) string {
	return fmt.Sprintf("\n%s\ntitle: %s\n%s\n", delim, title, delim)
}

// LinkCheckerPlugin is a custom plugin that checks whether links are internal or external
func LinkCheckerPlugin(baseDomain string) md.Plugin {
	return func(c *md.Converter) []md.Rule {
		return []md.Rule{
			{
				Filter: []string{"a"},
				Replacement: func(content string, node *goquery.Selection, opt *md.Options) *string {
					href, exists := node.Attr("href")
					if !exists {
						return nil
					}

					// Check if the link is external or internal
					if util.IsExternalUrl(href, baseDomain) {

						if strings.HasPrefix(href, "mailto:") {
							//Only do the replace if there is a mailto: prefix
							content = fmt.Sprintf("[%s](%s)", content, strings.Replace(href, "%40", "@", 1))
						} else {
							content = fmt.Sprintf("[%s](%s)", content, href)
						}

					} else {
						content = fmt.Sprintf("[%s](%s)", content, processInternalRef(href))
					}

					return &content
				},
			},
			{
				Filter: []string{"img"},
				Replacement: func(content string, node *goquery.Selection, opt *md.Options) *string {
					src, exists := node.Attr("src")
					if !exists {
						return nil
					}

					path := fmt.Sprintf("{{%%baseurl%%}}/%s", filepath.Join("images", filepath.Base(src)))
					content = fmt.Sprintf("![Image](%s)", path)

					return &content
				},
			},
			{
				Filter: []string{"video"},
				Replacement: func(content string, node *goquery.Selection, opt *md.Options) *string {
					src, exists := node.Attr("src")
					if !exists {
						return nil
					}

					path := fmt.Sprintf("{{%%baseurl%%}}/%s", filepath.Join("videos", filepath.Base(src)))
					content = fmt.Sprintf("{{< video %s >}}", path)

					return &content
				},
			},
		}
	}
}

/*removes html and htmls extensions in param string*/
func removeHtmlExtensions(href string) string {
	return strings.NewReplacer(
		".html", "",
		".htmls", "",
	).Replace(href)
}

/*adds the ref style*/
func processInternalRef(href string) string {
	cleanLink := removeHtmlExtensions(href)

	return fmt.Sprintf("{{< ref \"%s\" >}}", cleanLink)

}
