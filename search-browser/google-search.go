package searchbrowser

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

// SearchGoogle sends a search query to Google and returns a list of URLs from the search results.
func SearchGoogle(query string) ([]string, error) {
	// Base URL for Google search
	baseURL := "https://www.google.com/search"

	// Add query parameters
	params := url.Values{}
	params.Add("q", query)
	searchURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// Create an HTTP GET request
	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		return nil, err
	}

	// Set User-Agent to avoid being blocked by Google
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Go-http-client/1.1)")

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse the HTML response
	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}

	// Extract URLs from the HTML document
	var urls []string
	var extractLinks func(*html.Node)

	extractLinks = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					href := attr.Val

					// Filter out non-relevant links
					if strings.HasPrefix(href, "/url?q=") {
						// Extract the actual URL from the query parameter
						parsedURL := strings.Split(href, "&")[0]
						cleanURL := strings.TrimPrefix(parsedURL, "/url?q=")

						// Skip unwanted links (maps, accounts, support, etc.)
						if !strings.Contains(cleanURL, "maps.google.com") &&
							!strings.Contains(cleanURL, "support.google.com") &&
							!strings.Contains(cleanURL, "accounts.google.com") {
							urls = append(urls, cleanURL)
						}
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractLinks(c)
		}
	}

	// Start parsing from the root node
	extractLinks(doc)

	// Format the output with site names and URLs
	var formattedOutput []string
	for _, u := range urls {
		parsedURL, err := url.Parse(u)
		if err == nil {
			host := parsedURL.Hostname()
			formattedOutput = append(formattedOutput, fmt.Sprintf("NAME: %s, URL: %s\n", host, u))
		}
	}

	return formattedOutput, nil
}
