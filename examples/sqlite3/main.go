package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"paulwizviz/lotterystat/internal/euro"
	"paulwizviz/lotterystat/internal/sqlite"
	"time"

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
	db, err := sqlite.ConnectFile(dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = sqlite.InitalizeDB(db)
	if err != nil {
		log.Fatal(err)
	}

	stmt1, err := db.Prepare(euro.SQLiteInsertEuroStr)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt1.Close()

	e := euro.Draw{
		DrawDate:  time.Now(),
		DayOfWeek: time.Now().Weekday(),
		Ball1:     1,
		Ball2:     2,
		Ball3:     3,
		Ball4:     4,
		Ball5:     5,
		LS1:       1,
		LS2:       2,
		UKMarker:  "uk marker",
		DrawNo:    1,
	}

	r, err := sqlite.InsertDraw(context.TODO(), stmt1, e)
	if err != nil {
		log.Fatal("-->", e)
	}
	id, err := r.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rw, err := r.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Insert ID: %d Row affected: %d\n", id, rw)

}
