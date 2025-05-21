package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	name string
	id   uint32
}

type database struct {
	sql      *sql.DB
	buffer   []User
	inserter *sql.Stmt
}

func createDb(path string) (*database, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("open error: %w", err) 
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS test (id INTEGER PRIMARY KEY, name VARCHAR(32))")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	stmt, err := db.Prepare("insert into test (id, name) values (?, ?)")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	res := database{
		sql:      db,
		inserter: stmt,
		buffer:   make([]User, 0, 3),
	}

	return &res, nil
}

func (db *database) flush() error {
	transaction, err := db.sql.Begin()
	if err != nil {
		return err
	}

	for _, usr := range db.buffer {
		_, err := transaction.Stmt(db.inserter).Exec(usr.id, usr.name)
		if err != nil {
			transaction.Rollback()
			return err
		}
	}

	db.buffer = db.buffer[:0]
	return transaction.Commit()
}

func (db *database) insert(for_ins User) error {
	if len(db.buffer) == cap(db.buffer) {
		return errors.New("")
	}

	db.buffer = append(db.buffer, for_ins)
	if len(db.buffer) == cap(db.buffer) {
		err := db.flush()
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *database) get_from_id(id uint32) (*User, error) {
	const query = "select id, name from test where id = ?"

	var res User
	err := db.sql.QueryRow(query, id).Scan(&res.id, &res.name)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (db *database) close() error {
	defer func() {
		db.inserter.Close()
		db.sql.Close()
	}()

	err := db.flush()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	db, err := createDb("test.db")
	if err != nil {
		panic(err)
	}

	test := User{
		id:   0,
		name: "Meow",
	}

	err = db.insert(test)
	if err != nil {
		panic(err)
	}

	err = db.flush()
	if err != nil {
		panic(err)
	}

	res, rerr := db.get_from_id(0)
	if rerr != nil {
		panic(rerr)
	}

	fmt.Printf("%d %s", res.id, res.name)

	err = db.close()
	if err != nil {
		panic(err)
	}

}
