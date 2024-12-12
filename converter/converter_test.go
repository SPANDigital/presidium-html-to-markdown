package converter

import (
	"github.com/stretchr/testify/assert"
	"htmltomarkdown/config"
	"htmltomarkdown/models"
	"htmltomarkdown/test"
	"os"
	"testing"
	"testing/fstest"
)

func TestConvert(t *testing.T) {
	var dataPath = test.DataPath(t)

	converter := NewConverter("/", config.Config{
		Html: config.HtmlConfig{
			HeaderTags: []string{"h1"},
			Selector:   ".article",
			Remove:     []string{".article-title .permalink"},
			Replace: []models.DocReplacement{
				{
					Match:   ".tooltips-term",
					Replace: "{{< tooltip >}}",
				},
			},
		},
		Markdown: config.MarkdownConfig{},
	})

	dstDir := test.TempDir(t)
	defer os.RemoveAll(dstDir)

	err := converter.Convert(dataPath, dstDir)
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
