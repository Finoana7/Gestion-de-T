package main

import (
	"io"
	"log"
	"net/http"
	"nofi/middleware"
	"nofi/routes"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	var mux = http.NewServeMux()

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	})
	mux.Handle("/user/", middleware.CorsMiddleware(routes.Userhandler))
	mux.HandleFunc("/sold", middleware.CorsMiddleware(routes.SoldHandler))
	mux.HandleFunc("/recette", middleware.CorsMiddleware(routes.RecetteHandler))
	mux.HandleFunc("/trans", middleware.CorsMiddleware(routes.TransactionHandler))

	log.Println("About to listen on PORT :2005")

	// Démarrer le serveur
	http.ListenAndServe(":2005", mux)
}
