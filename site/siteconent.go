package site

import (
	"fmt"
	"net/http"
	"strings"

	"go-browser/types"

	"golang.org/x/net/html"
)

// SiteContent fetches and processes content from a given site based on provided options.
func SiteContent(options types.SiteOptions) (string, error) {
	// Create HTTP request
	req, err := http.NewRequest("GET", options.URL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Accept-Language", options.Language)

	// Execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to access the site: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	// Parse HTML
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to parse HTML: %w", err)
	}

	// Extract and filter content
	content, err := extractContent(doc, options)
	if err != nil {
		return "", err
	}

	return cleanContent(content), nil
}

// extractContent traverses the HTML tree and extracts content based on the provided options.
func extractContent(n *html.Node, options types.SiteOptions) (string, error) {
	var content []string
	seen := make(map[string]bool) // Tracks already seen content to avoid duplicates

	// Recursive function to traverse the HTML nodes
	var traverse func(*html.Node)
	traverse = func(node *html.Node) {
		if node.Type == html.ElementNode { // Process only element nodes
			// Skip unwanted tags based on the filter option
			if options.Filter && isUnwantedTag(node.Data) {
				return
			}
			// Check if the element matches the specified criteria
			if options.Element == "" || node.Data == options.Element {
				text := extractText(node)
				if text != "" && !seen[text] { // Avoid duplicate content
					seen[text] = true
					if options.IncludeAttributes {
						// Include attributes in the extracted content
						attr := collectAttributes(node)
						content = append(content, fmt.Sprintf("<%s %s>%s</%s>", node.Data, attr, text, node.Data))
					} else {
						// Add only the text content
						content = append(content, text)
					}
				}
			}
		}
		// Traverse the child nodes
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			traverse(child)
		}
	}

	traverse(n)                             // Start traversal from the root node
	return strings.Join(content, "\n"), nil // Return extracted content as a single string
}

// isUnwantedTag checks if the tag should be skipped.
func isUnwantedTag(tag string) bool {
	unwantedTags := map[string]bool{
		"script": true, "style": true, "link": true, "meta": true,
	}
	return unwantedTags[tag]
}

// extractText recursively collects visible text from a node.
func extractText(n *html.Node) string {
	if n.Type == html.TextNode {
		return strings.TrimSpace(n.Data)
	}
	var result strings.Builder
	for child := n.FirstChild; child != nil; child = child.NextSibling {
		result.WriteString(extractText(child))
	}
	return strings.TrimSpace(result.String())
}

// collectAttributes returns a string representation of node attributes.
func collectAttributes(n *html.Node) string {
	var attrs []string
	for _, attr := range n.Attr {
		attrs = append(attrs, fmt.Sprintf(`%s="%s"`, attr.Key, attr.Val))
	}
	return strings.Join(attrs, " ")
}

// cleanContent filters and removes unnecessary lines from the content.
func cleanContent(content string) string {
	lines := strings.Split(content, "\n")
	var cleanedLines []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || containsUnwantedKeywords(line) {
			continue
		}
		cleanedLines = append(cleanedLines, line)
	}
	return strings.Join(cleanedLines, "\n")
}

// containsUnwantedKeywords checks if a line contains unwanted keywords.
func containsUnwantedKeywords(line string) bool {
	unwantedKeywords := []string{"function", "envFlush", "__DEV__"}
	for _, keyword := range unwantedKeywords {
		if strings.Contains(line, keyword) {
			return true
		}
	}
	return false
}
