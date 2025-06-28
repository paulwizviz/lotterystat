package sqlops

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
)

const (
	UnspecifiedType SQLType = iota
	SQLiteType
	PSQLType
)

// SQLType represents a variant of a SQL base
// on a DB type
type SQLType int

func (s SQLType) String() string {
	switch s {
	case SQLiteType:
		return "SQLite"
	case PSQLType:
		return "PostgreSQL"
	default:
		return "Unspecified"
	}
}

// DriverType returns the SQL type,
func DriverType(db *sql.DB) SQLType {
	switch reflect.TypeOf(db.Driver()).String() {
	case "*sqlite3.SQLiteDriver":
		return SQLiteType
	case "*pq.Driver":
		return PSQLType
	default:
		return UnspecifiedType
	}
}

var (
	// Errors
	ErrCreateTbl   = errors.New("unable to create table")
	ErrPrepareStmt = errors.New("prepare statement")
	ErrWriteTbl    = errors.New("unable to write to table")
	ErrQueryTbl    = errors.New("query error")
	ErrDBConn      = errors.New("connection error")
)

// NewSQLiteMem instantiate a connection to SQLite
func NewSQLiteMem() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, fmt.Errorf("%w-%v", ErrDBConn, err)
	}
	return db, nil
}

// NewSQLiteFile instantiate a file based SQLite
func NewSQLiteFile(f string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", f)
	if err != nil {
		return nil, fmt.Errorf("%w-%v", ErrDBConn, err)
	}
	return db, nil
}

// NewPSQL instantiate a connection to PostgreSQL server
func NewPSQL(username string, password string, host string, port int, dbname string) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("%w-%v", ErrDBConn, err.Error())
	}
	return db, nil
}
