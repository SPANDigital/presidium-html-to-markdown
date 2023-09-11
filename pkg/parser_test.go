package pkg

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestParseArticles(t *testing.T) {
	var testdata = filepath.Join("../testdata", t.Name())
	var expected []Article
	mustUnmarshal(t, filepath.Join(testdata, "articles.json"), &expected)

	content := mustReadFileAsString(t, filepath.Join(testdata, "page.md"))
	actual, err := ParseArticles(content)
	assert.NoError(t, err, "failed to parse articles")
	assert.Equal(t, expected, actual)
}
