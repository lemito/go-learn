package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/lemito/web_and_bd/internal/db"
)

func main() {
	test_db, err := db.CreateDb("meow.db")
	if err != nil {
		test_db.Close()
		panic(err)
	}

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Pong!!")
	})

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {

		id_str := r.URL.Query().Get("id")
		if id_str == "" {
			http.Error(w, "", http.StatusBadRequest)
		}

		id, err := strconv.Atoi(id_str)
		if err != nil {
			http.Error(w, "", http.StatusBadRequest)
		}

		res, err := test_db.Get_from_id(uint32(id))
		if err != nil {
			http.Error(w, "", http.StatusBadRequest)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
