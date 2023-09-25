package repo

import (
	"database/sql"
	"errors"
	"paulwizviz/lotterystat/internal/euro"
	"paulwizviz/lotterystat/internal/sforl"
)

var (
	ErrCreateTbl   = errors.New("unable to create table")
	ErrPrepareStmt = errors.New("prepare statement")
	ErrWriteTbl    = errors.New("unable to write to table")
)

type drawType interface {
	euro.Draw | sforl.Draw
}

type structTag struct {
	FieldName string
	Tag       string
}

type CreateTableFn func(*sql.DB, interface{}) error
type WriteTableFn func(*sql.DB, interface{}) error
type WriteTxTableFn func(*sql.Tx, interface{}) error

var (
	CreateTable  CreateTableFn  = sqliteCreateTbl
	WriteTable   WriteTableFn   = sqliteWriteTbl
	WriteTxTable WriteTxTableFn = sqliteWriteTxTbl
)
