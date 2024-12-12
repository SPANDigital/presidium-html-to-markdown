package util

import (
	"fmt"
	"net/url"
	"path/filepath"
	"regexp"
	"strings"

	rake "github.com/afjoseph/RAKE.Go"
)

var (
	slugRe        = regexp.MustCompile("[^a-zA-Z0-9]+")
	specialCharRe = regexp.MustCompile(`[^\w\s\d]+`)
)

func ReplaceTemplate(text string, args []string) string {
	for i, arg := range args {
		placeholder := fmt.Sprintf("$%d", i+1)
		text = strings.Replace(text, placeholder, arg, -1)
	}
	return text
}

func Slugify(s string) string {
	val := slugRe.ReplaceAllString(s, "-")
	val = strings.Trim(val, "-")
	return strings.ToLower(val)
}

func IsURL(path string) bool {
	parsedURL, err := url.Parse(path)
	if err != nil {
		return false
	}
	return len(parsedURL.Scheme) > 0 && len(parsedURL.Host) > 0
}

func IsExternalUrl(path string, baseDomain string) bool {

	//Handle mailto: links
	if strings.HasPrefix(path, "mailto:") {
		return true
	}

	parsedURL, err := url.Parse(path)
	if err != nil {
		return false // Handle invalid URLs as internal by default
	}

	// Check if the hostname is different or if it's a relative link
	if parsedURL.Host != "" && !strings.Contains(parsedURL.Host, baseDomain) {
		return true
	}

	return false
}

func PathByType(contentType, filePath, assetDir string) string {
	filename := filepath.Base(filePath)
	if strings.HasPrefix(contentType, "image/") {
		return filepath.Join(assetDir, "/images", filename)
	} else if strings.HasPrefix(contentType, "video/") {
		return filepath.Join(assetDir, "/videos", filename)
	} else if strings.HasPrefix(contentType, "application/") {
		return filepath.Join(assetDir, "/files", filename)
	}
	return filePath
}

func FileNameWithoutExt(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func EscapeTitle(title string) string {
	if specialCharRe.MatchString(title) {
		title = strings.ReplaceAll(title, "\"", "")
		return fmt.Sprintf("\"%s\"", title)
	}
	return title
}

func FilenameFromTitle(title string) string {
	if len(title) > 30 {
		title = ExtractKeywords(title, 3)
	}
	return Slugify(title)
}

func ExtractKeywords(text string, count int) string {
	text = specialCharRe.ReplaceAllString(text, "")
	text = strings.ToLower(text)
	candidates := rake.RunRake(text)
	var words []string
	for _, candidate := range candidates {
		words = append(words, strings.Fields(candidate.Key)...)
	}

	lim := minInt(len(words), count)
	return trimNonKeywords(text, words[:lim])
}

func trimNonKeywords(text string, wordList []string) string {
	words := rake.SeperateWords(text)
	wordSet := make(map[string]struct{})
	for _, word := range wordList {
		wordSet[word] = struct{}{}
	}

	var newWords []string
	for _, word := range words {
		if _, exists := wordSet[word]; exists {
			newWords = append(newWords, word)
		}
	}
	return strings.Join(newWords, " ")
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
