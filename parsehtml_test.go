package main

import "testing"

func TestGetH1FromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputHTML string
		expected  string
	}{
		{
			name:      "only h1",
			inputHTML: "<h1>An H1 Header</h1>",
			expected:  "An H1 Header",
		},
		{
			name:      "basic HTML",
			inputHTML: "<html><body><h1>Test Title</h1></body></html>",
			expected:  "Test Title",
		},
		{
			name:      "basic HTML (multiple H1 headers)",
			inputHTML: "<html><body><h1>Test Title 1</h1><h1>Test Title 2</h1><h1>Test Title 3</h1></body></html>",
			expected:  "Test Title 1",
		},
		{
			name:      "basic HTML (no H1 header)",
			inputHTML: "<html><body><h2>Test Title</h2></body></html>",
			expected:  "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getH1FromHTML(tc.inputHTML)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
			}
			if actual != tc.expected {
				t.Errorf("Test %v - '%s' FAIL: expected: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})

	}
}
