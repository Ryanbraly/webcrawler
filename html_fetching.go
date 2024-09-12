package main

import (
	"fmt"
	"net/http"
	"io"
	"strings"
)

func getHTML(rawURL string) (string, error) {

	res, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("Trouble receiving data from URL: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return "", fmt.Errorf("Status code exceeds 400!  Code: ", res.StatusCode)
	}

	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("Content type is not HTML: %s", contentType)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("Trouble reading data from response body: %w", err)
	}

	return string(data), nil
}