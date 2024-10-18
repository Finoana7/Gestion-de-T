package routes

import (
	"net/http"
	"nofi/controller"
)

func SoldHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		controller.GetSold(w, r)
		return
	}

	http.Error(w, "", http.StatusMethodNotAllowed)
}
