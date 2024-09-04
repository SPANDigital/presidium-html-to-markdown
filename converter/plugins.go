package converter

import (
	"fmt"
	"htmltomarkdown/util"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
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

func handleExternalLinks() md.Plugin {
	return func(c *md.Converter) []md.Rule {
		return []md.Rule{
			{
				Filter: []string{"a"},
				Replacement: func(content string, selec *goquery.Selection, opt *md.Options) *string {

					nodeAttr := selec.Nodes[len(selec.Nodes)-1].Attr

					log.Debug(nodeAttr)

					for i := 0; i < len(nodeAttr); i++ {
						log.Debug(nodeAttr[i])

						key := nodeAttr[i].Key

						log.Debug(key)

						if key == "href" {

							href := nodeAttr[i]

							log.Debug("href ", href)
							log.Debug("href val ", href.Val)

							link := href.Val
							log.Debug("link ", link)

							content = "[" + content + "](" + link + ")"
						}

					}

					log.Debug("content ", content)
					return &content
				},
			},
		}
	}
}
