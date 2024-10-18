package routes

import (
	"net/http"
	"nofi/controller"
)

func RecetteHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		controller.GetRecette(w, r)
	case http.MethodPost:
		controller.AddRecette(w, r)
	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
	}
}
