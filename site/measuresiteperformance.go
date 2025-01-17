package site

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// MeasureSitePerformance checks site performance with or without live monitoring.
func MeasureSitePerformance(url string, timeout time.Duration, live bool) error {
	client := http.Client{
		Timeout: timeout,
	}

	performRequest := func() error {
		start := time.Now()
		resp, err := client.Get(url)
		if err != nil {
			return fmt.Errorf("failed to access the site: %w", err)
		}
		defer resp.Body.Close()

		_, err = io.Copy(io.Discard, resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body: %w", err)
		}

		elapsed := time.Since(start)
		fmt.Printf("Site: %s\n", url)
		fmt.Printf("Status Code: %d\n", resp.StatusCode)
		fmt.Printf("Load Time: %.2fs\n", elapsed.Seconds())
		fmt.Printf("Timestamp: %s\n\n", time.Now().Format("15:04:05"))
		return nil
	}

	if live {
		fmt.Println("Live monitoring started... (Press CTRL+C to stop)")
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if err := performRequest(); err != nil {
					fmt.Println("Error:", err)
				}
			}
		}
	} else {
		return performRequest()
	}
}
