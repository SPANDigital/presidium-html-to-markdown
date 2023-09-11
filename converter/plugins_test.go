package converter

import (
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/stretchr/testify/assert"
	"htmltomarkdown/test"
	"path/filepath"
	"testing"
)

func TestArticlePlugin(t *testing.T) {
	conv := md.NewConverter("", true, nil)
	conv.Use(ArticlePlugin([]string{"h1", "h2"}, ";;;"))

	html := test.MustReadFileAsString(t, filepath.Join("../test/data", t.Name(), "page.html"))
	markdown := test.MustReadFileAsString(t, filepath.Join("../test/data", t.Name(), "page.md"))

	actual, err := conv.ConvertString(html)
	assert.NoError(t, err)
	assert.Equal(t, actual, markdown)
}
