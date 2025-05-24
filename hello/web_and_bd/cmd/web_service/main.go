package main

import (
	"log"
	"net/http"

	"github.com/lemito/web_and_bd/cmd/web_service/handlers"
	"github.com/lemito/web_and_bd/internal/db"
)

func main() {
	log.Print("Start at localhost:8080")
	test_db, err := db.CreateDb("meow.db")
	if err != nil {
		test_db.Close()
		panic(err)
	}

	defer test_db.Close()

	h := handlers.CreateHandler(test_db)

	http.HandleFunc("/ping", handlers.PingHandler)
	http.HandleFunc("/get", h.GetByIdHandler)
	http.HandleFunc("/get_all", h.GetAllHandler)
	http.HandleFunc("/add_elem", h.AddElemHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
