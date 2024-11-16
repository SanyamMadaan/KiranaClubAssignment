package processors

import (
	"errors"
	"math/rand"
	"net/http"
	"store-image-processor/models"
	"time"
)

func ProcessImage(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	// Simulate GPU processing delay
	time.Sleep(time.Duration(rand.Intn(300)+100) * time.Millisecond)

	// Mock perimeter calculation
	width, height := 100, 200 // Mock dimensions
	perimeter := 2 * (width + height)
	return perimeter, nil
}

func ProcessJob(storeID string, imageURLs []string) (string, error) {
	// Fetch store info dynamically
	store, err := models.FetchStore(storeID)
	if err != nil {
		return "", errors.New("store not found: " + storeID)
	}

	// Process each image
	for _, url := range imageURLs {
		_, err := ProcessImage(url)
		if err != nil {
			return "", errors.New("failed to process image for store: " + store.StoreName)
		}
	}

	return "success", nil
}
