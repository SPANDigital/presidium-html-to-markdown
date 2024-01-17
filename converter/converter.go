package converter

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"htmltomarkdown/config"
	"htmltomarkdown/models"
	"htmltomarkdown/parser"
	"htmltomarkdown/util"
	"os"
	"path/filepath"
)

var ErrNoContentFound = errors.New("no content found")

type Converter struct {
	baseUrl string
	cfg     config.Config
}

func NewConverter(baseUrl string, cfg config.Config) *Converter {
	return &Converter{
		baseUrl: baseUrl,
		cfg:     cfg,
	}
}

func (c *Converter) Convert(src string, dst string) error {
	err := os.MkdirAll(dst, os.ModePerm)
	if err != nil {
		return err
	}

	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		isHtmlFile, err := filepath.Match("*.html", info.Name())
		if err != nil {
			return err
		}

		if !isHtmlFile {
			return nil
		}

		relSrc, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		relSrc = filepath.Dir(relSrc)
		return c.convertFile(info.Name(), path, relSrc, dst)
	})
}

func (c *Converter) convertFile(filename, src, relDir, dst string) error {
	content, err := c.readFile(src)
	if err != nil {
		if errors.Is(ErrNoContentFound, err) {
			log.Warnf("No content found, skipped page: %s", src)
			return nil
		}
		return err
	}

	markdown := HtmlConverter(c.baseUrl, relDir, c.cfg).Convert(content)
	articles, err := parser.ParseArticles(markdown, ";;;")
	if err != nil {
		return fmt.Errorf("%s: %s", err, src)
	}

	dst = filepath.Join(dst, relDir)
	path := filepath.Join(dst, util.FileNameWithoutExt(filename))
	if filename == "index.html" {
		path = dst
	}

	log.Info("Converted: ", path)
	return c.createArticles(path, articles)
}

func (c *Converter) readFile(path string) (*goquery.Selection, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		return nil, err
	}

	content := doc.Find(c.cfg.Html.Selector)
	if len(content.Nodes) == 0 {
		return nil, ErrNoContentFound
	}

	return content, nil
}

func (c *Converter) createArticles(dir string, articles []models.Article) error {
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	for i, article := range articles {
		var filename = article.FileName()
		if i == 0 {
			filename = "_index.md"
		}

		article.FrontMatter.Weight = i + 1
		articlePath := filepath.Join(dir, filename)

		err := os.WriteFile(articlePath, []byte(article.String()), os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
