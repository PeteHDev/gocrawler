package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("error: could not parse URL <%s>: %v\n", rawBaseURL, err)
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("error: could not parse URL <%s>: %v\n", rawCurrentURL, err)
		return
	}

	if baseURL.Host != currentURL.Host {
		return
	}

	normCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error: could not normalize URL <%s>: %v\n", rawCurrentURL, err)
		return
	}

	_, ok := pages[normCurrentURL]
	if !ok {
		pages[normCurrentURL] = 1
	} else {
		pages[normCurrentURL]++
		return
	}

	fmt.Printf("crawling <%s>...\n", rawCurrentURL)
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error: failed to fetch HTML from <%s>: %v\n", rawCurrentURL, err)
		return
	}

	urls, err := getURLsFromHTML(html, currentURL)
	if err != nil {
		fmt.Printf("error: failed to retrieve URLs from <%s>: %v\n", rawCurrentURL, err)
		return
	}

	for _, link := range urls {
		crawlPage(rawBaseURL, link, pages)
	}
}
