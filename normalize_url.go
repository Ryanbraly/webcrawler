package main

import (
		"net/url"
		"fmt"
		"strings"
		)

func normalizeURL(rawURL string) (string, error) {
	
	// Take raw URL and parse it into a URL struct
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("Problem parsing raw URL: %w", err)
	}

	// Use the Host and Path fields of the URL struct to create a nice, trimmed-down path
	path := parsedURL.Host + parsedURL.Path

	// Normalize capitalization and remove any trailing "/"s
	path = strings.ToLower(path)
	path = strings.TrimSuffix(path, "/")
	
	return (path), nil
}