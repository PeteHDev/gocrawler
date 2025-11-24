package main

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getH1FromHTML(html string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Printf("error: %v\n", err)
		return "", err
	}
	header := doc.Find("h1").First().Text()
	return header, nil
}

func getParagraphFromHTML(html string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Printf("error: %v\n", err)
		return "", err
	}
	var paragraph string
	paragraph = doc.Find("main").First().Find("p").First().Text()
	if paragraph == "" {
		paragraph = doc.Find("p").First().Text()
	}
	return paragraph, nil
}
