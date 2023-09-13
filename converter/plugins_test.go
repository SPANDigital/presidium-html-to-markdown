package converter

import (
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/stretchr/testify/assert"
	"htmltomarkdown/test"
	"path/filepath"
	"testing"
)

func TestArticlePlugin(t *testing.T) {
	var dataPath = test.DataPath(t)

	converter := md.NewConverter("", true, nil)
	converter.Use(ArticlePlugin([]string{"h1", "h2"}, ";;;"))

	html := test.MustReadFileAsString(t, filepath.Join(dataPath, "page.html"))
	markdown := test.MustReadFileAsString(t, filepath.Join(dataPath, "page.md"))

	actual, err := converter.ConvertString(html)
	assert.NoError(t, err)
	assert.Equal(t, actual, markdown)
}
