package handlers

import (
	"net/http"

	"github.com/PuncharatLiu/google-maps-scraper/internal/pkg/cref"
	"github.com/PuncharatLiu/google-maps-scraper/internal/pkg/db"
)

func CrefHandler(w http.ResponseWriter, r *http.Request) {
	cref.GetAllName(w, r)
	cref.CleanUp()
	inBoardStatus := cref.Cref(extractData.Titles)

	db.CreateBusiness(extractData, inBoardStatus)
}
