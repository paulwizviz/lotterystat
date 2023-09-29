// Package sqlite contains operations to persists data in
// SQLite DB.
//
// The driver for this package is https://github.com/mattn/go-sqlite3
package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"paulwizviz/lotterystat/internal/euro"
	"paulwizviz/lotterystat/internal/sforl"
)

var (
	ErrCreateTbl   = errors.New("unable to create table")
	ErrPrepareStmt = errors.New("prepare statement")
	ErrWriteTbl    = errors.New("unable to write to table")
	ErrConn        = errors.New("connection error")
)

func ConnectMem() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, fmt.Errorf("%w-%v", ErrConn, err)
	}
	return db, nil
}

func ConnectFile(f string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", f)
	if err != nil {
		return nil, fmt.Errorf("%w-%v", ErrConn, err)
	}
	return db, nil
}

func InitalizeDB(db *sql.DB) error {
	_, err := db.Exec(euro.SQLiteCreateTblStr)
	if err != nil {
		return fmt.Errorf("%w-%s", ErrCreateTbl, err.Error())
	}
	_, err = db.Exec(sforl.SQLiteCreateTblStr)
	if err != nil {
		return fmt.Errorf("%w-%s", ErrCreateTbl, err.Error())
	}
	return nil
}

func InsertDraw(ctx context.Context, stmt *sql.Stmt, draw any) (sql.Result, error) {
	switch d := draw.(type) {
	case euro.Draw:
		result, err := stmt.Exec(d.DrawDate, d.DayOfWeek, d.Ball1, d.Ball2, d.Ball3, d.Ball4, d.Ball5, d.LS1, d.LS2, d.UKMarker, d.DrawNo)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, nil
	}
}
