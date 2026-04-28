package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var PrewarmSecret = generateRandomSecret()

func generateRandomSecret() string {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(bytes)
}

// StartPrewarmer initializes a background goroutine that fetches all endpoints periodically
// to ensure the GlobalCache is always populated.
func StartPrewarmer(endpoints []string) {
	// Execute immediately in a goroutine
	go func() {
		// Initial wait to let the server start
		time.Sleep(2 * time.Second)

		for {
			log.Println("[Prewarmer] Starting batch fetch for", len(endpoints), "endpoints...")
			startTime := time.Now()

			for _, endpoint := range endpoints {
				url := fmt.Sprintf("http://127.0.0.1:8080%s", endpoint)
				
				req, err := http.NewRequest("GET", url, nil)
				if err != nil {
					log.Printf("[Prewarmer] Error creating request for %s: %v\n", url, err)
					continue
				}

				// Inject secret to bypass AuthMiddleware
				req.Header.Set("X-Prewarm", PrewarmSecret)

				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					log.Printf("[Prewarmer] Error fetching %s: %v\n", url, err)
					continue
				}

				// We must read and close the body to reuse connections
				_, _ = io.ReadAll(resp.Body)
				resp.Body.Close()
				
				// Small delay to prevent spiking CPU too aggressively
				time.Sleep(1 * time.Second)
			}

			log.Printf("[Prewarmer] Finished batch fetch in %v. Waiting for next cycle.\n", time.Since(startTime))

			// Run every 10 minutes
			time.Sleep(10 * time.Minute)
		}
	}()
}
