package main

import (
	"log"
	"os"
	"fmt"
	"net/http"
)


func main() {
	fmt.Println("starting your http server!")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "homepage")
	})
	http.ListenAndServe(":"+port, nil)
}
