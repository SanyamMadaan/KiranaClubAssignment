package main

import (
	"fmt"
	"log"
	"store-image-processor/models"
)

func main() {
	// Load store data from a CSV file
	err := models.LoadStoreMaster("store_master.csv") // Replace with your file path
	if err != nil {
		log.Fatalf("Failed to load store master data: %v", err)
	}

	fmt.Println("Store data loaded successfully.")
	// Add the rest of your application logic (e.g., API setup)
}
