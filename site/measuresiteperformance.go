package site

import (
	"fmt"
	"go-browser/types"
	"io"
	"log"
	"net/http"
	"time"
)

// SitePerformanceChecker handles site performance checks with or without live monitoring.
type SitePerformanceChecker struct {
	Client types.HTTPClient
	URL    string
}

// NewSitePerformanceChecker initializes a new SitePerformanceChecker with the given HTTP client and URL.
func NewSitePerformanceChecker(client types.HTTPClient, url string) *SitePerformanceChecker {
	return &SitePerformanceChecker{
		Client: client,
		URL:    url,
	}
}

// PerformRequest checks the performance of a single request and prints the results.
func (s *SitePerformanceChecker) PerformRequest() error {
	start := time.Now()
	resp, err := s.Client.Get(s.URL)
	if err != nil {
		return fmt.Errorf("failed to access the site: %w", err)
	}
	defer resp.Body.Close()

	// Discard response body content to measure performance without processing.
	_, err = io.Copy(io.Discard, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Measure elapsed time and print the results.
	elapsed := time.Since(start)
	fmt.Printf("Site: %s\n", s.URL)
	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Load Time: %.2fs\n", elapsed.Seconds())
	fmt.Printf("Timestamp: %s\n\n", time.Now().Format("15:04:05"))
	return nil
}

// Monitor continuously checks the site performance at a fixed interval.
func (s *SitePerformanceChecker) Monitor(interval time.Duration) {
	fmt.Println("Live monitoring started... (Press CTRL+C to stop)")
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		if err := s.PerformRequest(); err != nil {
			log.Println("Error:", err)
		}
	}
}

// MeasureSitePerformance initiates the site performance check with or without live monitoring.
func MeasureSitePerformance(url string, timeout time.Duration, live bool, interval int) error {
	client := &types.DefaultHTTPClient{
		Client: http.Client{Timeout: timeout},
	}

	checker := NewSitePerformanceChecker(client, url)

	if live {
		// Run live monitoring with a fixed interval of 5 seconds.
		checker.Monitor(time.Duration(interval) * time.Second)
		return nil
	}

	// Perform a single request check.
	return checker.PerformRequest()
}
