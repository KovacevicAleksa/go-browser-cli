package site

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// SiteContent fetches clean HTML content from a given URL with advanced filtering.
func SiteContent(url, element string, lang string, includeAttributes, filter bool) (string, error) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Accept-Language", lang)

	// Perform the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to access the site: %w", err)
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	// Parse the HTML
	parsedHTML, err := html.Parse(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to parse HTML: %w", err)
	}

	// Filter and clean the content
	var extractedContent []string
	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode {
			// Skip unwanted tags
			if filter && (n.Data == "script" || n.Data == "style" || n.Data == "link" || n.Data == "meta") {
				return
			}

			// Match specified element
			if element == "" || n.Data == element {
				if includeAttributes {
					attributes := collectAttributes(n)
					content := extractText(n)
					extractedContent = append(extractedContent, fmt.Sprintf("<%s %s>%s</%s>", n.Data, attributes, content, n.Data))
				} else {
					content := extractText(n)
					extractedContent = append(extractedContent, content)
				}
			}
		}

		// Recursively traverse child nodes
		for child := n.FirstChild; child != nil; child = child.NextSibling {
			traverse(child)
		}
	}
	traverse(parsedHTML)

	// Clean up the extracted content
	cleanedContent := cleanContent(strings.Join(extractedContent, "\n"))

	// Return the cleaned content
	if cleanedContent == "" {
		return "No relevant content found.", nil
	}
	return cleanedContent, nil
}

// extractText collects all visible text from a node and its children.
func extractText(n *html.Node) string {
	if n.Type == html.TextNode {
		return strings.TrimSpace(n.Data)
	}

	var textContent []string
	for child := n.FirstChild; child != nil; child = child.NextSibling {
		textContent = append(textContent, extractText(child))
	}
	return strings.Join(textContent, " ")
}

// collectAttributes returns a string representation of attributes in a node.
func collectAttributes(n *html.Node) string {
	var attributes []string
	for _, attr := range n.Attr {
		attributes = append(attributes, fmt.Sprintf(`%s="%s"`, attr.Key, attr.Val))
	}
	return strings.Join(attributes, " ")
}

// cleanContent removes unnecessary data from the extracted content.
func cleanContent(content string) string {
	lines := strings.Split(content, "\n")
	var cleanedLines []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		// Skip empty lines or lines containing scripts or unwanted keywords
		if line == "" || strings.Contains(line, "function") || strings.Contains(line, "envFlush") || strings.Contains(line, "__DEV__") {
			continue
		}
		cleanedLines = append(cleanedLines, line)
	}
	return strings.Join(cleanedLines, "\n")
}
