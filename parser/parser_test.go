package parser

import (
	"github.com/stretchr/testify/assert"
	"htmltomarkdown/models"
	"htmltomarkdown/test"
	"path/filepath"
	"testing"
)

func TestParseArticles(t *testing.T) {
	var testdata = filepath.Join("../test/data", t.Name())
	var expected []models.Article
	test.MustUnmarshal(t, filepath.Join(testdata, "articles.json"), &expected)

	content := test.MustReadFileAsString(t, filepath.Join(testdata, "page.md"))
	actual, err := ParseArticles(content, ";;;")
	assert.NoError(t, err, "failed to parse articles")
	assert.Equal(t, expected, actual)
}
