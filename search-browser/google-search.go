package searchbrowser

import (
	"encoding/json"
	"fmt"
	"go-browser/types"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

const googleAPIBaseURL = "https://www.googleapis.com/customsearch/v1"

// loadEnv loads environment variables from a .env file.
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// SearchGoogle uses Google Custom Search JSON API to perform a search.
func SearchGoogle(query string) ([]string, error) {

	loadEnv()

	// Get API key and Custom Search Engine ID from environment variables
	apiKey := os.Getenv("API_KEY")
	searchEngineID := os.Getenv("CSE_ID")

	// Ensure API Key and CSE ID are loaded correctly
	if apiKey == "" || searchEngineID == "" {
		return nil, fmt.Errorf("missing API Key or CSE ID")
	}

	// Prepare URL parameters for the search query
	params := url.Values{}
	params.Add("key", apiKey)        // Add API key to query parameters
	params.Add("cx", searchEngineID) // Add Custom Search Engine ID to query parameters
	params.Add("q", query)           // Add the search query to parameters

	// Final URL for the API request
	searchURL := fmt.Sprintf("%s?%s", googleAPIBaseURL, params.Encode())

	// Create HTTP GET request
	resp, err := http.Get(searchURL)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check HTTP status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected HTTP status: %s", resp.Status)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Parse the JSON response
	var result types.SearchResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	// Process the search results
	var formattedOutput []string
	for _, item := range result.Items {
		// Format each result with title and URL
		formattedOutput = append(formattedOutput, fmt.Sprintf("NAME: %s, URL: %s\n", item.Title, item.Link))
	}

	return formattedOutput, nil // Return the formatted search results
}
