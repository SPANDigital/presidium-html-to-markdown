package models

import (
	"fmt"
	"htmltomarkdown/util"
)

type Article struct {
	FrontMatter FrontMatter
	Content     string
}

type FrontMatter struct {
	Title  string `yaml:"title"`
	Weight int    `yaml:"weight"`
}

func NewArticle(fm *FrontMatter) *Article {
	return &Article{
		FrontMatter: *fm,
	}
}

func (a *Article) FileName() string {
	return fmt.Sprintf("%s.md", util.FilenameFromTitle(a.FrontMatter.Title))
}

func (a *Article) WriteLine(s string) {
	a.Content += fmt.Sprintln(s)
}

func (a *Article) String() string {
	return fmt.Sprintf("---\ntitle: \"%s\"\nweight: %d\n---\n%s", a.FrontMatter.Title, a.FrontMatter.Weight, a.Content)
}
