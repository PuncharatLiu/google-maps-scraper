package models

type ExtractData struct {
	Titles  []string `json:"titles"`
	Address []string `json:"address"`
	Website []string `json:"website"`
	Phone   []string `json:"phone"`
	Cid     []string `josn:"cid"`
}
