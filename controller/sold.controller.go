package controller

import (
	"encoding/json"
	"net/http"
	"nofi/data"
)

func GetSold(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(data.Sold)
}
