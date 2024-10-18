package data

import (
	"encoding/json"
	"io"
	"net/http"
	"nofi/models"
)

func AddRecette(w http.ResponseWriter, r http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var loginReq models.TAddRecette
	err = json.Unmarshal(body, &loginReq)

	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}
