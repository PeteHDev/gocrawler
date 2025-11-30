package main

import (
	"fmt"
	"net/url"
)

type PageData struct {
	URL            string
	H1             string
	FirstParagraph string
	OutgoingLinks  []string
	ImageURLs      []string
}

func extractPageData(html, pageURL string) PageData {
	var pd PageData
	pd.URL = pageURL
	h1, err := getH1FromHTML(html)
	if err != nil {
		fmt.Printf("warning: failed to find H1 header (%v)", err)
	} else {
		pd.H1 = h1
	}

	paragraph, err := getParagraphFromHTML(html)
	if err != nil {
		fmt.Printf("warning: failed to find first paragraph (%v)", err)
	} else {
		pd.FirstParagraph = paragraph
	}

	pageURLStruct, err := url.Parse(pageURL)
	if err != nil {
		fmt.Printf("warning: failed to parse page URL (%v)", err)
	}

	links, err := getURLsFromHTML(html, pageURLStruct)
	if err != nil {
		fmt.Printf("warning: failed to find outgoing links (%v)", err)
	} else {
		pd.OutgoingLinks = links
	}

	images, err := getImagesFromHTML(html, pageURLStruct)
	if err != nil {
		fmt.Printf("warning: failed to find image URLs (%v)", err)
	} else {
		pd.ImageURLs = images
	}

	return pd
}
