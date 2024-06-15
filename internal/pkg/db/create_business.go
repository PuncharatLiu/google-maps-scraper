package db

import (
	"fmt"
	"log"

	"github.com/PuncharatLiu/google-maps-scraper/models"
)

func CreateBusiness(business models.ExtractData, inBoardStatus []bool) {
	db, err := Connection()
	if err != nil {
		log.Fatalf("Error connect")
	}

	query := `
		INSERT INTO business (in_board, title, phone, address, website, cid)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (title) DO NOTHING
	`

	for i, status := range inBoardStatus {
		if status {
			continue
		}

		_, err := db.Exec(query, inBoardStatus[i], business.Titles[i], business.Phone[i], business.Address[i], business.Website[i], business.Cid[i])
		if err != nil {
			log.Fatalf("Error create business to PostgreSQL: %q", err)
		}
	}

	fmt.Println("Complete create business to PostgreSQL")
}
