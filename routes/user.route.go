package routes

import (
	"net/http"
	"nofi/controller"
	"nofi/middleware"
	"strings"
)

func Userhandler(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	path := strings.TrimPrefix(r.URL.Path, "/user")

	switch {
	case path == "/":
		middleware.JwtGuard(controller.GetOne).ServeHTTP(w, r)
	case path == "/login" && r.Method == http.MethodPost:
		controller.Login(w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
