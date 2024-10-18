package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"nofi/data"
	"nofi/helper"
	"nofi/models"

	"github.com/google/uuid"
)

func AddRecette(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var dataAdd models.TAddRecette
	err = json.Unmarshal(body, &dataAdd)

	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	data.Sold = data.Sold + dataAdd.Amount
	var recette models.TRecette = models.TRecette{ID: uuid.NewString(), Date: helper.GetDate(), Label: dataAdd.Label, Amount: dataAdd.Amount}

	data.Recette = append(data.Recette, recette)

	json.NewEncoder(w).Encode(recette)
}

func GetRecette(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(data.Recette)
}
