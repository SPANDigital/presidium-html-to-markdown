package converter

import (
	"fmt"
	"htmltomarkdown/config"
	"htmltomarkdown/models"
	"htmltomarkdown/util"
	"strings"

	html2md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/JohannesKaufmann/html-to-markdown/plugin"
	"github.com/PuerkitoBio/goquery"
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

func HtmlConverter(baseUrl, path string, cfg config.Config) *html2md.Converter {
	conv := html2md.NewConverter(path, true, &html2md.Options{})

	conv.Before(remove(cfg.Html.Remove), replace(cfg.Html.Replace))
	conv.Use(plugin.Table(), ArticlePlugin(cfg.Html.HeaderTags, ";;;"))
	conv.Use(LinkCheckerPlugin(baseUrl))
	conv.AddRules(rules...)

	return conv
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
