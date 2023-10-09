package main

import (
	"context"
	"database/sql"
	"log"
	"paulwizviz/lotterystat/internal/config"
	"paulwizviz/lotterystat/internal/dbutil"
	"paulwizviz/lotterystat/internal/worker"

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
	err = worker.InitalizeDB(ctx, DB)
	if err != nil {
		log.Fatal(err)
	}
	err = worker.PersistsDraw(ctx, DB)
	if err != nil {
		log.Fatal(err)
	}
	Execute()
}
