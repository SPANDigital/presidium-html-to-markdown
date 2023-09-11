package pkg

import (
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestArticlePlugin(t *testing.T) {
	conv := md.NewConverter("", true, nil)
	conv.Use(ArticlePlugin([]string{"h1", "h2"}, ";;;"))

	html := mustReadFileAsString(t, filepath.Join("../testdata", t.Name(), "page.html"))
	markdown := mustReadFileAsString(t, filepath.Join("../testdata", t.Name(), "page.md"))

	actual, err := conv.ConvertString(html)
	assert.NoError(t, err)
	assert.Equal(t, actual, markdown)
}
