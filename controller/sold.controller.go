package controller

import (
	"net/http"
	"nofi/data"
)

func GetSold(w http.ResponseWriter, r *http.Request) float64 {
	return data.Sold
}
