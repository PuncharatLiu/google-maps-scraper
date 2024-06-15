package handlers

import (
	"net/http"

	"github.com/PuncharatLiu/google-maps-scraper/internal/pkg/cref"
)

func CrefHandler(w http.ResponseWriter, r *http.Request) {
	cref.GetAllName(w, r)
	cref.CleanUp()
	cref.Cref(extractData.Titles)
}
