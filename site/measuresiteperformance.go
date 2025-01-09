package site

import (
	"fmt"
	"net/http"
	"time"
)

// MeasureSitePerformance tests the loading time and status code of a website.
func MeasureSitePerformance(url string, timeout time.Duration) error {
	start := time.Now() // Start timer

	// Set up HTTP client with timeout
	client := http.Client{
		Timeout: timeout,
	}

	// Perform the GET request
	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("failed to access the site: %w", err)
	}
	defer resp.Body.Close()

	// Calculate the time it took to load the site
	elapsed := time.Since(start)

	// Output basic site information
	fmt.Printf("Site: %s\n", url)
	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Load Time: %s\n", elapsed)

	return nil
}
