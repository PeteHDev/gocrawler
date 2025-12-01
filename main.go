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

	pages := make(map[string]int)
	fmt.Printf("starting crawl of: %v\n", os.Args[1])
	crawlPage(os.Args[1], os.Args[1], pages)
	for link, count := range pages {
		fmt.Printf("%s -> %d\n", link, count)
	}
}
