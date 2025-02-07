package types

import "net/http"

// CommandHandler represents a command with its associated handler function and description.
type CommandHandler struct {
	Command     string // The command string, e.g. "/help"
	Handler     func() // The function that will handle the command
	Description string // A description of what the command does
}

// RequestBody represents the body of a request to an API, including model parameters and messages.
type RequestBody struct {
	Model       string              `json:"model"`       // The model being used for the request
	Messages    []map[string]string `json:"messages"`    // A list of messages in the request
	MaxTokens   int                 `json:"max_tokens"`  // The maximum number of tokens for the request
	Temperature float64             `json:"temperature"` // The temperature for controlling randomness in the response
}

// ResponseBody represents the response from an API, including choices with messages.
type ResponseBody struct {
	Choices []struct {
		Message map[string]string `json:"message"` // The message returned in the choice
	} `json:"choices"` // List of choices in the response
}

// SearchResult represents the structure of the search result from the API.
type SearchResult struct {
	Items []struct {
		Title string `json:"title"` // The title of the search result
		Link  string `json:"link"`  // The link to the search result
	} `json:"items"` // List of items in the search result
}

// HTTPClient defines an interface for making HTTP requests.
// This allows us to use a mock client for testing.
type HTTPClient interface {
	Get(url string) (*http.Response, error) // The Get method to perform an HTTP GET request
}

// DefaultHTTPClient wraps the standard http.Client for making HTTP requests.
type DefaultHTTPClient struct {
	Client http.Client // The standard HTTP client used for requests
}

// Get performs an HTTP GET request using the default HTTP client.
func (d *DefaultHTTPClient) Get(url string) (*http.Response, error) {
	return d.Client.Get(url) // Executes the GET request
}

// SiteOptions defines options for fetching and filtering site content.
type SiteOptions struct {
	URL               string // The URL of the site
	Element           string // The HTML element to target
	Language          string // The language of the content
	IncludeAttributes bool   // Whether to include HTML attributes in the content
	Filter            bool   // Whether to apply a filter to the content
}
