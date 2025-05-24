package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/lemito/web_and_bd/internal/db"
)

type Handler struct {
	db *db.Database
}

func CreateHandler(db *db.Database) *Handler {
	return &Handler{db: db}
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong!!")
}

func (h *Handler) GetByIdHandler(w http.ResponseWriter, r *http.Request) {
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

	res, err := h.db.Get_from_id(uint32(id))
	if err != nil {
		http.Error(w, fmt.Sprintf("No such id or other err: %v", err), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *Handler) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	res, err := h.db.Get_all()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *Handler) AddElemHandler(w http.ResponseWriter, r *http.Request) {
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
		h.db.Insert(usr)
		fmt.Fprintf(w, "Added %s!", name)
	}
}
