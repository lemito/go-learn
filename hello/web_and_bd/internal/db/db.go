package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Name string `json:"id"`
	Id   uint32 `json:"name"`
}

type Database struct {
	sql      *sql.DB
	buffer   []User
	inserter *sql.Stmt
}

func CreateDb(path string) (*Database, error) {
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

	res := Database{
		sql:      db,
		inserter: stmt,
		buffer:   make([]User, 0, 3),
	}

	return &res, nil
}

func (db *Database) Flush() error {
	transaction, err := db.sql.Begin()
	if err != nil {
		return err
	}

	for _, usr := range db.buffer {
		_, err := transaction.Stmt(db.inserter).Exec(usr.Id, usr.Name)
		if err != nil {
			transaction.Rollback()
			return err
		}
	}

	db.buffer = db.buffer[:0]
	return transaction.Commit()
}

func (db *Database) Insert(for_ins User) error {
	if len(db.buffer) == cap(db.buffer) {
		return errors.New("")
	}

	db.buffer = append(db.buffer, for_ins)
	if len(db.buffer) == cap(db.buffer) {
		err := db.Flush()
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *Database) Get_from_id(id uint32) (*User, error) {
	const query = "select id, name from test where id = ?"

	var res User
	err := db.sql.QueryRow(query, id).Scan(&res.Id, &res.Name)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (db *Database) Get_all() (*[]User, error) {
	const query = "select id, name from test"

	var res []User
	rows, err := db.sql.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var tmp User
		err := rows.Scan(&tmp.Id, &tmp.Name)
		if err != nil {
			return nil, fmt.Errorf("error: %v", err)
		}
		res = append(res, tmp)
	}

	return &res, nil
}

func (db *Database) Close() error {
	defer func() {
		db.inserter.Close()
		db.sql.Close()
	}()

	err := db.Flush()
	if err != nil {
		return err
	}

	return nil
}

// func main() {
// 	db, err := createDb("test.db")
// 	if err != nil {
// 		panic(err)
// 	}

// 	test := User{
// 		id:   0,
// 		name: "Meow",
// 	}

// 	err = db.insert(test)
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = db.flush()
// 	if err != nil {
// 		panic(err)
// 	}

// 	res, rerr := db.get_from_id(0)
// 	if rerr != nil {
// 		panic(rerr)
// 	}

// 	fmt.Printf("%d %s", res.id, res.name)

// 	err = db.close()
// 	if err != nil {
// 		panic(err)
// 	}

// }
