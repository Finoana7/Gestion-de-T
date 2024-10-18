package routes

import (
	"net/http"
	"nofi/controller"
	"nofi/middleware"
)

func RecetteHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		middleware.JwtGuard(controller.GetRecette).ServeHTTP(w, r)
	case http.MethodPost:
		middleware.JwtGuard(controller.AddRecette).ServeHTTP(w, r)
	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
	}
}
