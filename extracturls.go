package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	baseURLString := strings.TrimSuffix(baseURL.String(), "/")
	refs := doc.Find("a[href]")
	urls := make([]string, 0, refs.Length())
	refs.Each(func(_ int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		hrefURL, err := url.Parse(href)
		if err != nil {
			fmt.Printf("error: failed to parse url <%v>\n", href)
			return
		}

		if hrefURL.Host == baseURL.Host {
			urls = append(urls, href)
		} else if hrefURL.Host == "" && hrefURL.Path != "" {
			urls = append(urls, baseURLString+"/"+strings.TrimPrefix(hrefURL.Path, "/"))
		}
	})

	return urls, nil
}
