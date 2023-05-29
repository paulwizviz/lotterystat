package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/paulwizviz/lotterystat/internal/csvdata"

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

	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = csvdata.TblCreateHdl[csvdata.EuroDraw](db, csvdata.EuroDraw{}, csvdata.SQLiteCreateEuroDrawTbl)
	if err != nil {
		log.Fatal(err)
	}

	eds, err := csvdata.TblReadHdl(db, csvdata.SQLiteReadEuroDrawTbl)
	if err != nil {
		log.Fatal(err)
	}

	lastDrwNo := len(eds) - 1

	ed := csvdata.EuroDraw{
		DrawDate:   time.Now(),
		DayOfWeek:  time.Friday,
		Ball1:      0,
		Ball2:      1,
		Ball3:      3,
		Ball4:      4,
		Ball5:      5,
		LS1:        1,
		LS2:        2,
		UKMarker:   "abc",
		EuroMarker: "efg",
		DrawNo:     uint64(lastDrwNo + 1),
	}

	err = csvdata.TblWriterHdl(db, ed, csvdata.SQLiteWriteEuroDrawTbl)
	if errors.Is(csvdata.ErrDuplicateEntry, errors.Unwrap(err)) {
		log.Fatal(err)
	}

	eds, err = csvdata.TblReadHdl(db, csvdata.SQLiteReadEuroDrawTbl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(eds)
}
