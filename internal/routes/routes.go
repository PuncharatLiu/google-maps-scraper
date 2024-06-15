package routes

import (
	"net/http"

	"github.com/PuncharatLiu/google-maps-scraper/internal/handlers"
)

// This is route for scrape process include extract data
func ScrapeRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/scrape", handlers.ScrapHandler)
}

/*
This is route for run cross reference process
*/
func Cref(mux *http.ServeMux) {
	mux.HandleFunc("/cref", handlers.CrefHandler)
}
