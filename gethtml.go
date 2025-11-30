package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func getHTML(rawURL string) (string, error) {
	client := http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", fmt.Errorf("error getHTML: %v", err)
	}
	req.Header.Set("User-Agent", "BootCrawler/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode > 399 {
		return "", fmt.Errorf("error getHTML: HTTP status code %d, %s", resp.StatusCode, resp.Status)
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("error getHTML: Content-Type of response is not text/html")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error getHTML: failed to read respnse body: %v", err)
	}
	defer resp.Body.Close()

	html := string(body)
	return html, nil
}
