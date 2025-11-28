package main

import (
	"fmt"
	"net/url"
)

func main() {
	inputURL := "https://blog.boot.dev"
	htmlBody := `
<html>
	<body>
		<a href="https://blog.boot.dev"><span>Boot.dev</span></a>
		<a href="https://wikipedia.org"><span>Boot.dev</span></a>
		<a href="https://google.com"><span>Boot.dev</span></a>
		<a href="https://youtube.com"><span>Boot.dev</span></a>
		<a href="https://blog.boot.dev/haha"><span>Boot.dev</span></a>
		<a href="https://blog.boot.dev/xoxo"><span>Boot.dev</span></a>
		<a href="https://blog.boot.dev/heehee"><span>Boot.dev</span></a>
	</body>
</html>
`
	baseURL, err := url.Parse(inputURL)
	if err != nil {
		fmt.Printf("Cpaka! %v", err)
	}
	urls, err := getURLsFromHTML(htmlBody, baseURL)
	if err != nil {
		fmt.Printf("Cpaka222! %v", err)
	}

	fmt.Printf("He cpaka!!! %v %v", urls, len(urls))
}
