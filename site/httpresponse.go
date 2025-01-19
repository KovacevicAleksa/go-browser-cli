package site

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// HttpRequest sends an HTTP request
func HttpRequest(url, method string, body []byte) (string, int, error) {
	// Create a new HTTP request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return "", 0, fmt.Errorf("error creating request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Execute the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", 0, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", 0, fmt.Errorf("error reading response body: %v", err)
	}

	// Return response body, status code, and any error
	return string(respBody), resp.StatusCode, nil
}
