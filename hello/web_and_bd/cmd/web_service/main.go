package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

	http.HandleFunc("/ping", handlers.PingHandler)

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {

		id_str := r.URL.Query().Get("id")
		if id_str == "" {
			http.Error(w, "No id", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(id_str)
		if err != nil {
			http.Error(w, "Bad id", http.StatusBadRequest)
			return
		}

		res, err := test_db.Get_from_id(uint32(id))
		if err != nil {
			http.Error(w, fmt.Sprintf("No such id or other err: %v", err), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	})

	http.HandleFunc("/get_all", func(w http.ResponseWriter, r *http.Request) {
		res, err := test_db.Get_all()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	})

	http.HandleFunc("/add_elem", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprintf(w, `
        	<form method="%s">
                <input type="text" name="name" placeholder="Name">
                <button type="submit">мяу</button>
            </form>
        `, http.MethodPost)
		} else if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				fmt.Fprintf(w, "Error: %v", err)
				return
			}

			name := r.FormValue("name")
			var usr db.User
			usr.Name = name
			test_db.Insert(usr)
			fmt.Fprintf(w, "Added %s!", name)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
