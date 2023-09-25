package main

import (
	"database/sql"
	"log"
	"os"
	"path"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dbPath := path.Join(pwd, "tmp", "sqlite")
	_, err = os.Stat(dbPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(dbPath, 0777)
		if err != nil {
			log.Fatal(err)
		}
	}

	dbFile := path.Join(pwd, "tmp", "sqlite", "data.db")
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}
