package site

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// SiteOptions defines options for fetching and filtering site content.
type SiteOptions struct {
	URL               string
	Element           string
	Language          string
	IncludeAttributes bool
	Filter            bool
}

// SiteContent fetches and processes content from a given site based on provided options.
func SiteContent(options SiteOptions) (string, error) {
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

// extractContent traverses the HTML and extracts content based on options.
func extractContent(n *html.Node, options SiteOptions) (string, error) {
	var content []string
	var traverse func(*html.Node)
	traverse = func(node *html.Node) {
		if node.Type == html.ElementNode {
			// Skip unwanted tags if filtering is enabled
			if options.Filter && isUnwantedTag(node.Data) {
				return
			}

			// Match specific element or all elements if not specified
			if options.Element == "" || node.Data == options.Element {
				text := extractText(node)
				if options.IncludeAttributes {
					attr := collectAttributes(node)
					content = append(content, fmt.Sprintf("<%s %s>%s</%s>", node.Data, attr, text, node.Data))
				} else {
					content = append(content, text)
				}
			}
		}

		// Recurse into child nodes
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			traverse(child)
		}
	}
	traverse(n)

	return strings.Join(content, "\n"), nil
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
