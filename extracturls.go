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

	refs := doc.Find("a[href]")
	urls := make([]string, 0, refs.Length())
	refs.Each(func(_ int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		if strings.TrimSpace(href) == "" {
			return
		}

		hrefURL, err := url.Parse(href)
		if err != nil {
			fmt.Printf("error: failed to parse url <%v>\n", href)
			return
		}

		absolute := baseURL.ResolveReference(hrefURL)
		urls = append(urls, absolute.String())
	})

	return urls, nil
}

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	srcs := doc.Find("img[src]")
	images := make([]string, 0, srcs.Length())
	srcs.Each(func(_ int, s *goquery.Selection) {
		src, _ := s.Attr("src")
		if strings.TrimSpace(src) == "" {
			return
		}

		srcURL, err := url.Parse(src)
		if err != nil {
			fmt.Printf("error: failed to parse src <%v>\n", src)
			return
		}

		absolute := baseURL.ResolveReference(srcURL)
		images = append(images, absolute.String())
	})

	return images, nil
}
