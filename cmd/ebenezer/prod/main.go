package main

import (
	"context"
	"database/sql"
	"log"
	"paulwizviz/lotterystat/internal/config"
	"paulwizviz/lotterystat/internal/dbutil"

	_ "github.com/mattn/go-sqlite3"
)

var (
	DB *sql.DB
)

func main() {
	err := config.Initilalize()
	if err != nil {
		log.Fatal(err)
	}
	DB, err = dbutil.ConnectFile(config.DBConn())
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.TODO()
	err = initalizeDB(ctx, DB)
	if err != nil {
		log.Fatal(err)
	}
	err = persistsDraw(ctx, DB)
	if err != nil {
		log.Fatal(err)
	}
	Execute()
}
