package converter

import (
	"fmt"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"htmltomarkdown/util"
	"strings"
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
					content = strings.Replace(content, `#`, `\#`, -1)
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
