package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	maxConcurrency := 10
	cfg, err := configure(os.Args[1], maxConcurrency)
	if err != nil {
		fmt.Printf("error - configure: %v", err)
		os.Exit(1)
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(os.Args[1])
	cfg.wg.Wait()

	for page, data := range cfg.pages {
		fmt.Print("==============================\n==============================\n")
		fmt.Printf("<%s>\n", page)
		fmt.Printf("Title: %s\n", data.H1)
		fmt.Printf("Excerpt: %s\n", data.FirstParagraph)
		fmt.Printf("Outgoing links: %d\n", len(data.OutgoingLinks))
		fmt.Printf("Image URLs: %d\n", len(data.ImageURLs))
	}
}
