package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"paulwizviz/lotterystat/internal/dbutil"
	"paulwizviz/lotterystat/internal/euro"
	"paulwizviz/lotterystat/internal/sforl"

	_ "github.com/mattn/go-sqlite3"
)

func IntializeDB(ctx context.Context, db *sql.DB) {
	err := euro.CreateTable(ctx, db)
	if err != nil {
		log.Println(err)
	}
	err = sforl.CreateTable(ctx, db)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	ops := flag.String("ops", "", "db operations")
	flag.Parse()
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
	db, err := dbutil.ConnectFile(dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	ctx := context.TODO()
	if *ops == "initialize" {
		fmt.Println("Initializing DB ....")
		IntializeDB(ctx, db)
	} else if *ops == "populate" {
		fmt.Println("Populating DB ....")
	} else if *ops == "list" {
		fmt.Println("List all .....")
	} else {
		fmt.Println("No actions")
	}
}
