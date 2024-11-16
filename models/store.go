package models

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
)

type Store struct {
	AreaCode  string
	StoreName string
	StoreID   string
}

var Stores = map[string]Store{} // Dynamic store storage

// LoadStoreMaster loads store data from a CSV file into the Stores map
func LoadStoreMaster(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, _ = reader.Read() // Skip header row

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		store := Store{
			AreaCode:  record[0],
			StoreName: record[1],
			StoreID:   record[2],
		}
		Stores[store.StoreID] = store
	}
	return nil
}

// FetchStore dynamically looks up store details by StoreID
func FetchStore(storeID string) (*Store, error) {
	store, exists := Stores[storeID]
	if !exists {
		return nil, errors.New("store not found")
	}
	return &store, nil
}
