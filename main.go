package main

import (
	"io"
	"log"
	"net/http"
	"nofi/routes"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	var mux = http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	})
	mux.HandleFunc("/user/", routes.Userhandler)
	mux.HandleFunc("/sold", routes.SoldHandler)
	mux.HandleFunc("/recette", routes.RecetteHandler)
	mux.HandleFunc("/trans", routes.TransactionHandler)

	log.Println("About to listen on PORT :2005")

	err := http.ListenAndServe(":2005", mux)
	log.Fatal(err)
}
