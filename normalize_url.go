package main

import (
	"net/url"
	"strings"
)

func normalizeURL(urlStr string) (string, error) {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	// Get hostname and path
	host := strings.ToLower(parsedURL.Host)
	path := parsedURL.Path

	// Remove trailing slash if present
	if path != "/" && strings.HasSuffix(path, "/") {
		path = strings.TrimSuffix(path, "/")
	}

	// Combine host and path
	if path == "" {
		return host, nil
	}
	return host + path, nil
}
