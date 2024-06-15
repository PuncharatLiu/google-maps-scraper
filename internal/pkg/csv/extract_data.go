package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/PuncharatLiu/google-maps-scraper/models"
)

func ExtractData(resultFile *os.File) (models.ExtractData, error) {
	// Test
	fmt.Println("call ExtractData")

	reader := csv.NewReader(resultFile)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error read records: %v", err)
		return models.ExtractData{}, err
	}

	var extractData models.ExtractData
	for _, record := range records {
		if len(record) >= 2 {
			extractData.Titles = append(extractData.Titles, record[1])
			extractData.Address = append(extractData.Address, record[2])
			extractData.Website = append(extractData.Website, record[3])
			extractData.Phone = append(extractData.Phone, record[4])
			extractData.Cid = append(extractData.Cid, record[5])
		}
	}

	// Test
	fmt.Println("extractData: ", extractData)

	return extractData, nil
}
