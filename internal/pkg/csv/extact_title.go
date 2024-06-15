package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func ExtractTitle(resultFile *os.File) ([]string, error) {
	// Test
	fmt.Println("call ExtractTitle")

	// Create new reader.
	reader := csv.NewReader(resultFile)

	// Read all records from resultFile
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error read records: %v", err)
		return nil, err
	}

	// Extract title
	var titles []string
	for _, record := range records {
		if len(record) >= 2 {
			title := record[1]
			titles = append(titles, title)
		}
	}

	return titles, nil
}
