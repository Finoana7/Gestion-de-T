package routes

import (
	"net/http"
	"nofi/controller"
	"nofi/middleware"
)

func SoldHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		middleware.JwtGuard(controller.GetSold).ServeHTTP(w, r)
		return
	}

	http.Error(w, "", http.StatusMethodNotAllowed)
}
