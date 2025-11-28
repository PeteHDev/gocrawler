package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetURLsFromHTMLAbsolute(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody := `
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
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := getURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{
		"https://blog.boot.dev",
		"https://blog.boot.dev/haha",
		"https://blog.boot.dev/xoxo",
		"https://blog.boot.dev/heehee",
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestGetURLsFromHTMLRelative(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody := `
<html>
	<body>
		<a href="/some/path"><span>Boot.dev</span></a>
		<a href="https://wikipedia.org"><span>Boot.dev</span></a>
		<a href="https://google.com"><span>Boot.dev</span></a>
		<a href="https://youtube.com"><span>Boot.dev</span></a>
		<a href="haha"><span>Boot.dev</span></a>
		<a href="/xoxo"><span>Boot.dev</span></a>
		<a href="/heehee/"><span>Boot.dev</span></a>
	</body>
</html>
`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := getURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{
		"https://blog.boot.dev/some/path",
		"https://blog.boot.dev/haha",
		"https://blog.boot.dev/xoxo",
		"https://blog.boot.dev/heehee/",
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestGetURLsFromHTMLNone(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody := `
<html>
	<body>
		<a href="https://wikipedia.org"><span>Boot.dev</span></a>
		<a href="https://google.com"><span>Boot.dev</span></a>
		<a href="https://youtube.com"><span>Boot.dev</span></a>
	</body>
</html>
`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := getURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestGetURLsFromHTMLMixed(t *testing.T) {
	inputURL := "https://blog.boot.dev/"
	inputBody := `
<html>
	<body>
		<a href="/some/path"><span>Boot.dev</span></a>
		<a href="https://wikipedia.org"><span>Boot.dev</span></a>
		<a href="https://google.com"><span>Boot.dev</span></a>
		<a href="https://youtube.com"><span>Boot.dev</span></a>
		<a href="https://blog.boot.dev/haha"><span>Boot.dev</span></a>
		<a href="xoxo"><span>Boot.dev</span></a>
		<a href="/heehee/"><span>Boot.dev</span></a>
	</body>
</html>
`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := getURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{
		"https://blog.boot.dev/some/path",
		"https://blog.boot.dev/haha",
		"https://blog.boot.dev/xoxo",
		"https://blog.boot.dev/heehee/",
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
