package config

import "htmltomarkdown/models"

type Config struct {
	Html     HtmlConfig
	Markdown MarkdownConfig

	// List of domains to whitelist for links and images
	WhiteList []string

	// Directory to copy assets to
	AssetDir string

	// Directory to copy content to
	ContentDir string

	// Debug mode
	Debug bool
}

type HtmlConfig struct {
	// CSS-selector for the part of the page to extract and convert
	Selector string

	// List of header tags for article titles (h1, h2, h3 etc..)
	HeaderTags []string

	// List of CSS-selectors for elements to remove before conversion
	Remove []string

	// List of HTML replacements to perform before conversion
	Replace []models.DocReplacement
}

type MarkdownConfig struct {
	// List of Regex replacements to perform after conversion
	Replace []models.RegexReplace
}
