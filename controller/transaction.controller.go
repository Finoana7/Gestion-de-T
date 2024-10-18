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

func AddTRansaction(w http.ResponseWriter, r http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var dataAdd models.TAddTransaction
	err = json.Unmarshal(body, &dataAdd)

	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if data.Sold < dataAdd.Amount {
		http.Error(w, "Solde insufisant", http.StatusBadRequest)
		return
	}

	data.Sold = data.Sold - dataAdd.Amount
	transaction := models.TTransaction{ID: uuid.NewString(), Date: helper.GetDate(), Label: dataAdd.Label, Amount: dataAdd.Amount}

	data.Transaction = append(data.Transaction, transaction)

	json.NewEncoder(w).Encode(transaction)
}

func GetTransaction(w http.ResponseWriter, r http.Request) {
	json.NewEncoder(w).Encode(data.Transaction)
}
