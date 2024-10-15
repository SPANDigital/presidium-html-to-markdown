package converter

import (
	"fmt"
	"htmltomarkdown/util"
	"net/url"
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
					// content = strings.Replace(content, `#`, `\#`, -1)
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
					if isExternalLink(href, baseDomain) {
						content = fmt.Sprintf("[External Link: %s](%s)", content, href)
					} else {
						content = fmt.Sprintf("[Internal Link: %s](%s)", content, href)
					}

					return &content
				},
			},
		}
	}
}

// isExternalLink checks if the link is external based on the domain
func isExternalLink(link string, baseDomain string) bool {
	parsedURL, err := url.Parse(link)
	if err != nil {
		return false // Handle invalid URLs as internal by default
	}

	// Check if the hostname is different or if it's a relative link
	if parsedURL.Host != "" && !strings.Contains(parsedURL.Host, baseDomain) {
		return true
	}

	return false
}

// check that all external links are types or parsed correctly and are linking to external sites
// check that all internal links are parsed correctly with no confusing/extra chars
// check mailto, adirs etc, that they link to the correct application. (behavior) and if internal => what happens when we add the `({{<ref` style
// Maybe remove .html or whatever at the end of the link before applying the `({{<ref` style
// Track this in your ticket - track time spent and add comments on the ticket for visibility.
// Mpilo: to look at linking other file types and base url specific issues
// Mpilo: or the baseurl style: return fmt.Sprintf("{{%%baseurl%%}}/%s", path) for file links / images
