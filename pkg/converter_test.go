package pkg

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"
)

func TestConvert(t *testing.T) {
	var testdata = filepath.Join("../testdata", t.Name())
	c := NewConverter(Config{
		Html: HtmlConfig{
			HeaderTags: []string{"h1"},
			Selector:   ".article",
			Remove:     []string{".article-title .permalink"},
			Replace: []DocReplacement{
				{
					Match:   ".tooltips-term",
					Replace: "{{< tooltip >}}",
				},
			},
		},
		Markdown: MarkdownConfig{},
	})

	dstDir := testDir(t)
	defer os.RemoveAll(dstDir)

	err := c.Convert(testdata, dstDir)
	assert.NoError(t, err)

	err = fstest.TestFS(
		os.DirFS(dstDir),
		"_index.md",
		"authoring-workflow.md",
		"content-maintenance.md",
		"learning-objectives.md",
		"markdown.md",
		"menu-structure.md",
		"micro-articles.md",
		"templates.md",
	)
	assert.NoError(t, err)
}
