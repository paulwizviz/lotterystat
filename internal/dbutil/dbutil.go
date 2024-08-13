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

func SQLiteConnectMem() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, fmt.Errorf("%w-%v", ErrDBConn, err)
	}
	return db, nil
}

func SQLiteConnectFile(f string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", f)
	if err != nil {
		return nil, fmt.Errorf("%w-%v", ErrDBConn, err)
	}
	return db, nil
}

func PSQLConnect(username string, password string, host string, port int, dbname string) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("%w-%v", ErrDBConn, err.Error())
	}
	return db, nil
}
