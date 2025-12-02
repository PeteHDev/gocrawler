package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func writeCSVReport(pages map[string]PageData, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating csv report: %v", err)
	}
	writer := csv.NewWriter(file)
	err = writer.Write([]string{
		"page_url",
		"h1",
		"first_paragraph",
		"outgoing_link_urls",
		"image_urls",
	})
	if err != nil {
		return fmt.Errorf("error writing csv report: %v", err)
	}

	for _, data := range pages {
		outgoing_link_urls := strings.Join(data.OutgoingLinks, ";")
		image_urls := strings.Join(data.ImageURLs, ";")
		err = writer.Write([]string{
			data.URL,
			data.H1,
			data.FirstParagraph,
			outgoing_link_urls,
			image_urls,
		})
		if err != nil {
			return fmt.Errorf("error writing csv report: %v", err)
		}
	}

	return nil
}
