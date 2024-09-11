package converter

import (
	"fmt"
	"htmltomarkdown/util"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html"
)

const curlBrkOp = "{"
const curlBrkCl = "}"
const href = "href"
const ref = "ref"
const colon = "\""
const ltSign = "<"
const gtSign = ">"
const htmlEx = ".html"
const emptyStr = ""
const http = "http"
const https = "https"
const space = " "
const mailto = "mailto:"
const rdar = "rdar:"

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

					content = processAttr(content, nodeAttr)

					log.Debug("content ", content)
					return &content
				},
			},
		}
	}
}

func processInternalRef(link string) string {

	iniStr := curlBrkOp + curlBrkOp + ltSign + space + ref + space + colon
	endStr := colon + space + gtSign + curlBrkCl + curlBrkCl

	innerLink := iniStr + strings.Replace(link, htmlEx, emptyStr, -1) + endStr

	return innerLink

}

func processPrefix(content string, link string) string {

	if strings.HasPrefix(link, http) || strings.HasPrefix(link, https) || linkContainsReservedString(link) {
		content = "[" + content + "](" + link + ")"
	} else {
		content = "[" + content + "](" + processInternalRef(link) + ")"
	}

	return content

}

func linkContainsReservedString(link string) bool {

	contains := false

	for key := range getReservedStringsMap() {

		if strings.Contains(link, key) {
			contains = true
		}

	}

	return contains

}

func getReservedStringsMap() map[string]bool {

	reservedStr := map[string]bool{
		mailto: true,
		rdar:   true,
	}

	return reservedStr

}

func processAttr(content string, nodeAttr []html.Attribute) string {

	for i := 0; i < len(nodeAttr); i++ {
		log.Debug(nodeAttr[i])

		key := nodeAttr[i].Key

		log.Debug(key)

		if key == href {

			href := nodeAttr[i]

			log.Debug("href ", href)
			log.Debug("href val ", href.Val)

			link := href.Val
			log.Debug("link ", link)

			content = processPrefix(content, link)

		}

	}

	return content

}
