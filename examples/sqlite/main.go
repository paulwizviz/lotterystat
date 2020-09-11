package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dbDir := path.Join(wd, "tmp", "sqlite")
	err = os.MkdirAll(dbDir, 0755)
	if err != nil {
		log.Fatal(err)
	}
	dbfile := path.Join(dbDir, "data.db")
	fmt.Println(dbfile)

	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec()
	stmt, err = db.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec("Nic", "Raboy")
	rows, err := db.Query("SELECT id, firstname, lastname FROM people")
	if err != nil {
		log.Fatal(err)
	}
	var id int
	var firstname string
	var lastname string
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	}
}
