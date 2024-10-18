package routes

import (
	"net/http"
	"nofi/controller"
	"nofi/middleware"
)

func TransactionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		middleware.JwtGuard(controller.GetTransaction).ServeHTTP(w, r)
	case http.MethodPost:
		middleware.JwtGuard(controller.AddTRansaction).ServeHTTP(w, r)
	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
	}
}
