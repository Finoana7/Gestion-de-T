package main

import (
	"io"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	var mux = http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	})
	// mux.HandleFunc("/user/", routes.Userhandler)

	log.Println("About to listen on PORT :2005")

	err := http.ListenAndServe(":2005", mux)
	log.Fatal(err)
}
