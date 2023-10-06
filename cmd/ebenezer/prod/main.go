package main

import (
	"context"
	"database/sql"
	"log"
	"paulwizviz/lotterystat/internal/config"
	"paulwizviz/lotterystat/internal/dbutil"
	"paulwizviz/lotterystat/internal/euro"
	"paulwizviz/lotterystat/internal/sforl"

	_ "github.com/mattn/go-sqlite3"
)

var (
	DB *sql.DB
)

func initalizeDB(ctx context.Context, db *sql.DB) error {
	err := euro.CreateTable(ctx, db)
	if err != nil {
		return err
	}
	err = sforl.CreateTable(ctx, db)
	if err != nil {
		return err
	}
	return nil
}

func persistsDraw(ctx context.Context, db *sql.DB) error {
	err := euro.PersistsCSV(ctx, db, 3)
	if err != nil {
		return err
	}
	err = sforl.PersistsCSV(ctx, db, 3)
	if err != nil {
		return err
	}
	return nil
}

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
