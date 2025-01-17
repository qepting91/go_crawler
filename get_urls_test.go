package main

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:      "basic absolute URL",
			inputURL:  "https://blog.boot.dev",
			inputBody: `<html><body><a href="https://blog.boot.dev/path">Link</a></body></html>`,
			expected:  []string{"https://blog.boot.dev/path"},
		},
		{
			name:      "relative URL",
			inputURL:  "https://blog.boot.dev",
			inputBody: `<html><body><a href="/path">Link</a></body></html>`,
			expected:  []string{"https://blog.boot.dev/path"},
		},
		{
			name:     "multiple mixed URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
				<html><body>
					<a href="/relative/path">Link1</a>
					<a href="https://other.com/absolute">Link2</a>
					<a href="path">Link3</a>
				</body></html>`,
			expected: []string{
				"https://blog.boot.dev/relative/path",
				"https://other.com/absolute",
				"https://blog.boot.dev/path",
			},
		},
		{
			name:     "nested elements",
			inputURL: "https://blog.boot.dev",
			inputBody: `
				<html><body>
					<div>
						<a href="/path1"><span>Link1</span></a>
						<p><a href="/path2">Link2</a></p>
					</div>
				</body></html>`,
			expected: []string{
				"https://blog.boot.dev/path1",
				"https://blog.boot.dev/path2",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("%s: got unexpected error: %v", tc.name, err)
				return
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("%s: expected %v, got %v", tc.name, tc.expected, actual)
			}
		})
	}
}
