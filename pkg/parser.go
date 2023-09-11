package pkg

import (
	"bufio"
	"bytes"
	"gopkg.in/yaml.v2"
	"io"
	"strings"
)

const (
	articleDelimiter = ";;;"
)

// ParseArticles takes a markdown file and splits it into a list of Article
// It expects the articles to be delimited with the articleDelimiter
func ParseArticles(content string) ([]Article, error) {
	var articles []Article
	scanner := bufio.NewScanner(strings.NewReader(content))
	scanner.Scan()
	for {
		article, err := parseArticle(scanner)
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

func parseArticle(scanner *bufio.Scanner) (*Article, error) {
	frontMatter, err := parseFrontMatter(scanner)
	if err != nil {
		return nil, err
	}

	var article = NewArticle(frontMatter)
	for scanner.Scan() {
		line := scanner.Text()
		if scanner.Text() == articleDelimiter { // start of next article
			return article, err
		}
		article.WriteLine(line)
	}

	return article, io.EOF
}

func parseFrontMatter(scanner *bufio.Scanner) (*FrontMatter, error) {
	var buf bytes.Buffer
	for scanner.Scan() {
		line := scanner.Text()
		if line == articleDelimiter {
			break
		}
		buf.Write(scanner.Bytes())
	}

	var frontMatter FrontMatter
	if err := yaml.Unmarshal(buf.Bytes(), &frontMatter); err != nil {
		return nil, err
	}
	return &frontMatter, nil
}
