package converter

import (
	"fmt"
	"htmltomarkdown/util"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
)

// ArticlePlugin converts HTML headers to Hugo article headers in Markdown.
// E.g., '# This is a header' becomes '---\ntitle: This is a header\n---'.
func ArticlePlugin(filters []string, delim string) md.Plugin {
	return func(c *md.Converter) []md.Rule {
		return []md.Rule{
			{
				Filter: filters,
				Replacement: func(content string, selection *goquery.Selection, opt *md.Options) *string {
					if strings.TrimSpace(content) == "" {
						return nil
					}

					// Clean up and escape the content
					content = strings.ReplaceAll(content, "\n", " ")
					content = strings.ReplaceAll(content, "\r", " ")
					content = strings.ReplaceAll(content, `#`, `\#`)
					content = strings.TrimSpace(content)

					// Checks if the content is inside a link (anchor tag).
					if selection.ParentsFiltered("a").Length() > 0 {
						text := opt.StrongDelimiter + content + opt.StrongDelimiter
						text = md.AddSpaceIfNessesary(selection, text)
						return &text
					}

					// Generate the article header
					title := util.EscapeTitle(selection.Text())
					header := fmt.Sprintf("\n%s\ntitle: %s\n%s\n", delim, title, delim)
					return &header
				},
			},
		}
	}
}

// Plugin to handle nested unordered and ordered lists in Markdown.
func handleNestedLists() md.Plugin {
	return func(c *md.Converter) []md.Rule {
		return []md.Rule{
			{
				Filter: []string{"ul", "ol"}, // Target <ul> and <ol> tags.
				Replacement: func(content string, selection *goquery.Selection, opt *md.Options) *string {
					var markdown strings.Builder
					processListItems(&markdown, selection, 0, selection.Is("ol"))
					result := markdown.String()
					return &result
				},
			},
		}
	}
}

// Processes nested list items recursively, handling both ordered and unordered lists.
func processListItems(markdown *strings.Builder, selection *goquery.Selection, level int, isOrdered bool) {
	selection.Children().Each(func(i int, s *goquery.Selection) {
		prefix := getListPrefix(level, i, isOrdered)

		// Get and clean list item content
		itemContent := cleanListItemContent(s.Text())

		// Add the list item to the markdown output
		markdown.WriteString(prefix + itemContent + "\n")

		// Recursively process nested lists
		s.Find("ol, ul").Each(func(_ int, nested *goquery.Selection) {
			processListItems(markdown, nested, level+1, nested.Is("ol"))
		})
	})
}

// Returns the prefix for the list item, accounting for ordered/unordered lists.
func getListPrefix(level, index int, isOrdered bool) string {
	indent := strings.Repeat("    ", level)
	if isOrdered {
		return fmt.Sprintf("%s%d. ", indent, index+1)
	}
	return indent + "- "
}

// Cleans up list item content by removing extra newlines and trimming spaces.
func cleanListItemContent(content string) string {
	content = strings.ReplaceAll(content, "\n", " ")
	return strings.Join(strings.Fields(content), " ")
}

// Plugin to ignore links inside code blocks and return plain text.
func ignoreLinksInCodeBlocks() md.Plugin {
	return func(c *md.Converter) []md.Rule {
		return []md.Rule{
			{
				Filter: []string{"a"}, // Target <a> tags (links).
				Replacement: func(content string, selection *goquery.Selection, opt *md.Options) *string {
					// Checks if the selection is inside a code block.
					if selection.ParentsFiltered("code").Length() > 0 {
						linkText := selection.Text()
						return &linkText // Return plain text of the link if inside a code block.
					}
					return nil
				},
			},
		}
	}
}
