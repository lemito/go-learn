package main

import (
	"github.com/lemito/web_and_bd/internal/db"
)

func main() {
	test_db, err := db.CreateDb("meow.db")
}