package dbutil

import (
	"database/sql"
	"errors"
	"fmt"
)

var (
	ErrDBCreateTbl   = errors.New("unable to create table")
	ErrDBPrepareStmt = errors.New("prepare statement")
	ErrDBInsertTbl   = errors.New("unable to write to table")
	ErrDBQueryTbl    = errors.New("query error")
	ErrDBConn        = errors.New("connection error")
)

func ConnectMem() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, fmt.Errorf("%w-%v", ErrDBConn, err)
	}
	return db, nil
}

func ConnectFile(f string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", f)
	if err != nil {
		return nil, fmt.Errorf("%w-%v", ErrDBConn, err)
	}
	return db, nil
}
