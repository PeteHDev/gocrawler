package main

import (
	"reflect"
	"testing"
)

func TestExtractPageData(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody := `
<html>
	<body>
        <h1>Test Title</h1>
        <p>This is the first paragraph.</p>
        <a href="/link1">Link 1</a>
        <img src="/image1.jpg" alt="Image 1">
		<a href="https://blog.boot.dev/link2">Link 2</a>
		<img src="https://somewebsite.com/image2.jpg" alt="Image 2">
    </body>
	<main>
		<p>This is the second paragraph but it is in main.</p>
	</main>
</html>
`

	actual := extractPageData(inputBody, inputURL)

	expected := PageData{
		URL:            "https://blog.boot.dev",
		H1:             "Test Title",
		FirstParagraph: "This is the second paragraph but it is in main.",
		OutgoingLinks: []string{
			"https://blog.boot.dev/link1",
			"https://blog.boot.dev/link2",
		},
		ImageURLs: []string{
			"https://blog.boot.dev/image1.jpg",
			"https://somewebsite.com/image2.jpg",
		},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\nexpected %+v\ngot      %+v", expected, actual)
	}
}

func TestExtractPageDataMissingFields(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputHTML string
		expected  PageData
	}{
		{
			name:     "No Title",
			inputURL: "https://blog.boot.dev",
			inputHTML: `
<html>
	<body>
        <p>This is the first paragraph.</p>
        <a href="/link1">Link 1</a>
        <img src="/image1.jpg" alt="Image 1">
		<a href="https://blog.boot.dev/link2">Link 2</a>
		<img src="https://somewebsite.com/image2.jpg" alt="Image 2">
    </body>
	<main>
		<p>This is the second paragraph but it is in main.</p>
	</main>
</html>
`,
			expected: PageData{
				URL:            "https://blog.boot.dev",
				H1:             "",
				FirstParagraph: "This is the second paragraph but it is in main.",
				OutgoingLinks: []string{
					"https://blog.boot.dev/link1",
					"https://blog.boot.dev/link2",
				},
				ImageURLs: []string{
					"https://blog.boot.dev/image1.jpg",
					"https://somewebsite.com/image2.jpg",
				},
			},
		},
		{
			name:     "First Paragraph NOT In Main",
			inputURL: "https://blog.boot.dev",
			inputHTML: `
<html>
	<body>
		<h1>Test Title</h1>
        <p>This is the first paragraph.</p>
        <a href="/link1">Link 1</a>
        <img src="/image1.jpg" alt="Image 1">
		<a href="https://blog.boot.dev/link2">Link 2</a>
		<img src="https://somewebsite.com/image2.jpg" alt="Image 2">
    </body>
</html>
`,
			expected: PageData{
				URL:            "https://blog.boot.dev",
				H1:             "Test Title",
				FirstParagraph: "This is the first paragraph.",
				OutgoingLinks: []string{
					"https://blog.boot.dev/link1",
					"https://blog.boot.dev/link2",
				},
				ImageURLs: []string{
					"https://blog.boot.dev/image1.jpg",
					"https://somewebsite.com/image2.jpg",
				},
			},
		},
		{
			name:     "First Paragraph In Main",
			inputURL: "https://blog.boot.dev",
			inputHTML: `
<html>
	<body>
		<h1>Test Title</h1>
        <p>This is the first paragraph.</p>
        <a href="/link1">Link 1</a>
        <img src="/image1.jpg" alt="Image 1">
		<a href="https://blog.boot.dev/link2">Link 2</a>
		<img src="https://somewebsite.com/image2.jpg" alt="Image 2">
    </body>
	<main>
		<p>This is the second paragraph but it is in main.</p>
	</main>
</html>
`,
			expected: PageData{
				URL:            "https://blog.boot.dev",
				H1:             "Test Title",
				FirstParagraph: "This is the second paragraph but it is in main.",
				OutgoingLinks: []string{
					"https://blog.boot.dev/link1",
					"https://blog.boot.dev/link2",
				},
				ImageURLs: []string{
					"https://blog.boot.dev/image1.jpg",
					"https://somewebsite.com/image2.jpg",
				},
			},
		},
		{
			name:     "NO First Paragraph",
			inputURL: "https://blog.boot.dev",
			inputHTML: `
<html>
	<body>
		<h1>Test Title</h1>
        <a href="/link1">Link 1</a>
        <img src="/image1.jpg" alt="Image 1">
		<a href="https://blog.boot.dev/link2">Link 2</a>
		<img src="https://somewebsite.com/image2.jpg" alt="Image 2">
    </body>
</html>
`,
			expected: PageData{
				URL:            "https://blog.boot.dev",
				H1:             "Test Title",
				FirstParagraph: "",
				OutgoingLinks: []string{
					"https://blog.boot.dev/link1",
					"https://blog.boot.dev/link2",
				},
				ImageURLs: []string{
					"https://blog.boot.dev/image1.jpg",
					"https://somewebsite.com/image2.jpg",
				},
			},
		},
		{
			name:     "NO Outgoing Links",
			inputURL: "https://blog.boot.dev",
			inputHTML: `
<html>
	<body>
		<h1>Test Title</h1>
        <p>This is the first paragraph.</p>
        <img src="/image1.jpg" alt="Image 1">
		<img src="https://somewebsite.com/image2.jpg" alt="Image 2">
    </body>
	<main>
		<p>This is the second paragraph but it is in main.</p>
	</main>
</html>
`,
			expected: PageData{
				URL:            "https://blog.boot.dev",
				H1:             "Test Title",
				FirstParagraph: "This is the second paragraph but it is in main.",
				OutgoingLinks:  []string{},
				ImageURLs: []string{
					"https://blog.boot.dev/image1.jpg",
					"https://somewebsite.com/image2.jpg",
				},
			},
		},
		{
			name:     "NO Image URLs",
			inputURL: "https://blog.boot.dev",
			inputHTML: `
<html>
	<body>
		<h1>Test Title</h1>
        <p>This is the first paragraph.</p>
        <a href="/link1">Link 1</a>
		<a href="https://blog.boot.dev/link2">Link 2</a>
    </body>
	<main>
		<p>This is the second paragraph but it is in main.</p>
	</main>
</html>
`,
			expected: PageData{
				URL:            "https://blog.boot.dev",
				H1:             "Test Title",
				FirstParagraph: "This is the second paragraph but it is in main.",
				OutgoingLinks: []string{
					"https://blog.boot.dev/link1",
					"https://blog.boot.dev/link2",
				},
				ImageURLs: []string{},
			},
		},
		{
			name:     "NO Fields of Interest",
			inputURL: "https://blog.boot.dev",
			inputHTML: `
<html>
	<body>
			<div><b>Outrageous!!!</b></div>
    </body>
</html>
`,
			expected: PageData{
				URL:            "https://blog.boot.dev",
				H1:             "",
				FirstParagraph: "",
				OutgoingLinks:  []string{},
				ImageURLs:      []string{},
			},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := extractPageData(tc.inputHTML, tc.inputURL)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("\nTest %d: %s\nexpected %+v\ngot      %+v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
