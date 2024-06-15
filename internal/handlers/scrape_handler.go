package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/PuncharatLiu/google-maps-scraper/internal/pkg/csv"
	"github.com/PuncharatLiu/google-maps-scraper/models"
)

type ScraperRequest struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

var extractData models.ExtractData

func ScrapHandler(w http.ResponseWriter, r *http.Request) {
	var req ScraperRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// handle empty string
	if req.Key == "" || req.Name == "" {
		log.Println("Key or Name is empty")
		return
	}

	// createQueryFile(req.Key)
	csv.CreateQueryFile(req.Key)

	// Set the command and arguments
	cmd := exec.Command("../../google_map_scraper/google-maps-scraper", "-input", "../../test_output/query.txt", "-results", req.Name, "-exit-on-inactivity", "1m", "-depth", "1")

	// Run the command
	if err := cmd.Run(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Open the generated CSV file
	csvFile, err := os.Open(req.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer csvFile.Close()

	data, err := csv.ExtractData(csvFile)
	if err != nil {
		log.Fatalf("Error extract titles: %v", err)
	}

	extractData = data

	// Test
	fmt.Println("ExtractData: ", extractData)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(extractData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Complete scraping")
}
