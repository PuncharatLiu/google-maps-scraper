package routes

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"os/exec"
)

type ScraperRequest struct {
	InputFile  string `json:"inputFile"`
	ResultFile string `json:"resultFile"`
}

func RunScraper(w http.ResponseWriter, r *http.Request) {
	// Parse JSON input
	var req ScraperRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set default filenames if not provided
	if req.InputFile == "" {
		req.InputFile = "example-queries.txt"
	}
	if req.ResultFile == "" {
		req.ResultFile = "restaurants-in-cyprus.csv"
	}

	// Set the command and arguments
	cmd := exec.Command("./google-maps-scraper", "-input", req.InputFile, "-results", req.ResultFile, "-exit-on-inactivity", "3m")

	// Run the command
	if err := cmd.Run(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Open the generated CSV file
	csvFile, err := os.Open(req.ResultFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer csvFile.Close()

	// Set the headers for the file download
	w.Header().Set("Content-Disposition", "attachment; filename="+req.ResultFile)
	w.Header().Set("Content-Type", "text/csv")

	// Copy the file content to the response
	if _, err := io.Copy(w, csvFile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// func init() {
// 	http.HandleFunc("/run-scraper", runScraper)

// 	// Start the server
// 	log.Println("Server started on :8080")
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		log.Fatal(err)
// 	}
// }
