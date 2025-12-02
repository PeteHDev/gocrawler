package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	maxConcurrency := getMaxConcurrency()
	maxPages := getMaxPages()

	cfg, err := configure(os.Args[1], maxConcurrency, maxPages)
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

	writeCSVReport(cfg.pages, "report.csv")
}

func getMaxConcurrency() int {
	if len(os.Args) <= 2 {
		return 3
	}

	maxConcurrency, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("invalid <max concurrency> value. expecting base-10 integer.")
		os.Exit(1)
	}
	return maxConcurrency
}

func getMaxPages() int {
	if len(os.Args) <= 3 {
		return 100
	}

	maxPages, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Printf("invalid <max pages> value. expecting base-10 integer.")
		os.Exit(1)
	}
	return maxPages
}
