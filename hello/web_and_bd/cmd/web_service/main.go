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
	log.Print("Start")
	test_db, err := db.CreateDb("meow.db")
	if err != nil {
		test_db.Close()
		panic(err)
	}

	defer test_db.Close()

	test := db.User{
		Id:   0,
		Name: "Meow",
	}

	err = test_db.Insert(test)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Pong!!")
	})

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {

		id_str := r.URL.Query().Get("id")
		if id_str == "" {
			http.Error(w, "", http.StatusBadRequest)
			return

		}

		id, err := strconv.Atoi(id_str)
		if err != nil {
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		res, err := test_db.Get_from_id(uint32(id))
		if err != nil {
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
