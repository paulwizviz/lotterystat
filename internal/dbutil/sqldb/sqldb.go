// Package sqldb implements operations to manipulate SQL databases
package sqldb

import (
	"database/sql"
	"errors"
	"reflect"

	"github.com/paulwizviz/lotterystat/internal/dbutil"
)

var (
	ErrUnableToCreateTable = errors.New("unable to create table")
)

// TblCreator is an interface that wrap an operation to create Table
type TblCreator[T dbutil.CSVDataType] interface {
	Create(*sql.DB, string, T) error
}

// TblCreateHdlFunc register a function to create an SQL Table to persists
// CSV data types
type TblCreateHdlFunc[T dbutil.CSVDataType] func(*sql.DB, string, T) error

func (t TblCreateHdlFunc[T]) Create(db *sql.DB, tblname string, data T) error {
	return t(db, tblname, data)
}

// TblCreateHdl is responsible for the execution of SQL Table creation table
// for a handler function
func TblCreateHdl[T dbutil.CSVDataType](db *sql.DB, data T, hdl TblCreator[T]) error {

	tblname := reflect.TypeOf(data).Name()

	err := hdl.Create(db, tblname, data)
	if err != nil {
		return err
	}
	return nil
}

// TblWriter is an interface that wrap an operation to write CSV Data Types
// to an SQL Table
type TblWriter[T dbutil.CSVDataType] interface {
	Write(*sql.DB, string, T) error
}

// TblWriterHdlFunc register a function to store CSV Data types to SQL Table
type TblWriterHdlFunc[T dbutil.CSVDataType] func(db *sql.DB, tblName string, data T) error

func (t TblWriterHdlFunc[T]) Write(db *sql.DB, tblname string, data T) error {
	return t(db, tblname, data)
}

// TblWriterHdl is responsible for the execution of operation to store
// CSV data type in SQL Table
func TblWriterHdl[T dbutil.CSVDataType](db *sql.DB, data T, hdl TblWriterHdlFunc[T]) error {

	tblname := reflect.TypeOf(data).Name()
	err := hdl.Write(db, tblname, data)
	if err != nil {
		return err
	}
	return nil
}
