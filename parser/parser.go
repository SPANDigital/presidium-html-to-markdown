package parser

import (
	"bufio"
	"bytes"
	"gopkg.in/yaml.v2"
	"htmltomarkdown/models"
	"io"
	"strings"
)

// ParseArticles takes a markdown file and splits it into a list of Article
// It expects the articles to be delimited with the articleDelimiter
func ParseArticles(content string, delim string) ([]models.Article, error) {
	var articles []models.Article
	scanner := bufio.NewScanner(strings.NewReader(content))
	scanner.Scan()
	for {
		article, err := parseArticle(scanner, delim)
		if err != nil {
			if err == io.EOF {
				articles = append(articles, *article)
				break
			}
			return nil, err
		}
		articles = append(articles, *article)
	}
	return articles, nil
}

func parseArticle(scanner *bufio.Scanner, delim string) (*models.Article, error) {
	frontMatter, err := parseFrontMatter(scanner, delim)
	if err != nil {
		return nil, err
	}

	var article = models.NewArticle(frontMatter)
	for scanner.Scan() {
		line := scanner.Text()
		if scanner.Text() == delim { // start of next article
			return article, err
		}
		article.WriteLine(line)
	}

	return article, io.EOF
}

func parseFrontMatter(scanner *bufio.Scanner, delim string) (*models.FrontMatter, error) {
	var buf bytes.Buffer
	for scanner.Scan() {
		line := scanner.Text()
		if line == delim {
			break
		}
		buf.Write(scanner.Bytes())
	}

	var frontMatter models.FrontMatter
	if err := yaml.Unmarshal(buf.Bytes(), &frontMatter); err != nil {
		return nil, err
	}
	return &frontMatter, nil
}
