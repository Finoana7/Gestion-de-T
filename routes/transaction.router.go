package routes

import (
	"net/http"
	"nofi/controller"
)

func TransactionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		controller.GetTransaction(w, r)
	case http.MethodPost:
		controller.AddTRansaction(w, r)
	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
	}
}
