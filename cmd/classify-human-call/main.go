package main

import (
	"log"
	"net/http"

	"github.com/PuncharatLiu/google-maps-scraper/internal/routes"
	"github.com/rs/cors"
)

func initServer() {
	mux := http.NewServeMux()
	routes.ScrapeRoutes(mux)
	routes.Cref(mux)

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler(mux)

	// Start the server
	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}

func main() {
	initServer()
}
