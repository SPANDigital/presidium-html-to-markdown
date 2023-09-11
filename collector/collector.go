package collector

import (
	"errors"
	"fmt"
	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
	"htmltomarkdown/util"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func Collect(baseURL string, dst string) error {
	log.Infof("Collecting: %s", baseURL)
	filter, err := compileUrlFilter(baseURL)
	if err != nil {
		return err
	}

	c := colly.NewCollector(
		colly.URLFilters(filter),
		colly.CacheDir("./cache"),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		absoluteURL := resolveURL(e.Request.URL, href)
		visit(c, absoluteURL)
	})

	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		src := e.Attr("src")
		absoluteURL := resolveURL(e.Request.URL, src)
		visit(c, absoluteURL)
	})

	c.OnResponse(func(res *colly.Response) {
		contentType := res.Headers.Get("Content-Type")
		path := util.PathByType(contentType, res.Request.URL.Path)
		path = filepath.Join(dst, path)
		ext := filepath.Ext(path)

		if strings.HasPrefix(contentType, "text/html") && len(ext) == 0 {
			path = filepath.Join(path, "index.html")
		}

		err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
		if err != nil {
			log.Errorf("Error creating directory: %s, error: %s", path, err)
			return
		}

		if err := res.Save(path); err != nil {
			log.Errorf("Error saving file: %s, error: %s", path, err)
			return
		}

		log.Infof("Collected: %s", res.Request.URL.Path)
	})

	return c.Visit(baseURL)
}

func visit(c *colly.Collector, url string) {
	if err := c.Visit(url); err != nil {
		if errors.Is(err, colly.ErrNoURLFiltersMatch) || errors.Is(err, colly.ErrAlreadyVisited) {
			return
		}
		log.Errorf("Error visiting url: %s, error: %s", url, err)
	}
}

func resolveURL(baseURL *url.URL, link string) string {
	ref, err := url.Parse(link)
	if err != nil {
		log.Printf("Error parsing link %s: %v", link, err)
		return ""
	}

	ref.Fragment = "" // remove fragment (#) to avoid duplicate visits ot the same page
	return baseURL.ResolveReference(ref).String()
}

func compileUrlFilter(url string) (*regexp.Regexp, error) {
	urlFilter := regexp.QuoteMeta(url)
	urlFilter = fmt.Sprintf("%s.*", urlFilter)
	return regexp.Compile(urlFilter)
}
