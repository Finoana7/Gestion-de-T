package routes

import (
	"net/http"
	"nofi/controller"
	"strings"
)

func SoldHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/sold")

	if path != "/" && r.Method == http.MethodGet {
		controller.GetSold(w, r)
		return
	}

	http.Error(w, "", http.StatusMethodNotAllowed)
}
