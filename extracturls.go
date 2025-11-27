package main

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	refs := doc.Find("a[href]")
	urls := make([]string, 0, refs.Length())
	refs.Each(func(_ int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		urls = append(urls, href)
	})

	return urls, nil
}
