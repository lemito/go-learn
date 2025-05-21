package main

import (
	"fmt"
	"log"
	"net/http"

	// "github.com/lemito/web_and_bd/internal/db"
)

func main() {
	// test_db, err := db.CreateDb("meow.db")

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Pong!!")
    })

	log.Fatal(http.ListenAndServe(":8080", nil))
}