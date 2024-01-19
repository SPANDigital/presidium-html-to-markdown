package parser

import (
	"github.com/stretchr/testify/assert"
	"htmltomarkdown/models"
	"htmltomarkdown/test"
	"path/filepath"
	"testing"
)

func TestParseArticles(t *testing.T) {
	var dataPath = test.DataPath(t)
	var expected []models.Article
	test.MustUnmarshal(t, filepath.Join(dataPath, "articles.json"), &expected)

	content := test.MustReadFileAsString(t, filepath.Join(dataPath, "page.md"))
	actual, err := ParseArticles(content, ";;;")
	assert.NoError(t, err, "failed to parse articles")
	assert.Equal(t, expected, actual)
}

func TestParseArticlesNoHeader(t *testing.T) {
	var dataPath = test.DataPath(t)
	var expected []models.Article
	test.MustUnmarshal(t, filepath.Join(dataPath, "articles.json"), &expected)

	content := test.MustReadFileAsString(t, filepath.Join(dataPath, "page.md"))
	actual, err := ParseArticles(content, ";;;")
	assert.NoError(t, err, "failed to parse articles")
	assert.Equal(t, expected, actual)
}
