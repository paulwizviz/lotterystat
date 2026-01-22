package sqlops

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "modernc.org/sqlite"
)

var (
	// Errors
	ErrCreateTbl     = errors.New("unable to create table")
	ErrCreateTxn     = errors.New("unable to create transaction")
	ErrCloseStmt     = errors.New("unable to close statment")
	ErrPrepareStmt   = errors.New("prepare statement")
	ErrExecuteQuery  = errors.New("execute query error")
	ErrExecuteWriter = errors.New("execute write error")
	ErrDBConn        = errors.New("connection error")
)

// NewSQLiteMem instantiate a connection to SQLite
func NewSQLiteMem() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "file::memory:?cache=shared")
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDBConn, err)
	}
	return db, nil
}

// NewSQLiteFile instantiate a file based SQLite
func NewSQLiteFile(f string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", f)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDBConn, err)
	}
	return db, nil
}

// TblCreator is a function type to help create
// db table
type TblCreator func(context.Context, *sql.Tx) error

func (t TblCreator) Create(ctx context.Context, tx *sql.Tx) error {
	return t(ctx, tx)
}

// CreateTables is a wrapper function to create Table based on the
// a slice of table creators
func CreateTables(ctx context.Context, db *sql.DB, creators ...TblCreator) error {
	tx, err := db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCreateTxn, err)
	}
	committed := false
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		} else if !committed {
			tx.Rollback()
		}
	}()

	for _, creator := range creators {
		if err = creator.Create(ctx, tx); err != nil {
			return fmt.Errorf("%w: %v", ErrCreateTbl, err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("%w:%w", ErrCreateTbl, err)
	}
	committed = true

	return nil
}

// RowWriter is a function type to support callback to write a row of data
type RowWriter func(context.Context, *sql.Stmt, any) error

func Writer(ctx context.Context, db *sql.DB, rawStmt string, dataList []any, rowWriter RowWriter) error {
	if len(dataList) == 0 {
		return nil
	}

	stmt, err := db.PrepareContext(ctx, rawStmt)
	if err != nil {
		return fmt.Errorf("%w:%w", ErrPrepareStmt, err)
	}
	defer stmt.Close()

	tx, err := db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCreateTxn, err)
	}
	for _, data := range dataList {
		err := rowWriter(ctx, stmt, data)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	tx.Commit()
	return nil
}

// QueryScanner is a function type to support callback to read a row of data
type QueryScanner func(*sql.Rows) (any, error)

func Query(ctx context.Context, db *sql.DB, scanner QueryScanner, rawQuery string, args ...any) ([]any, error) {
	stmt, err := db.PrepareContext(ctx, rawQuery)
	if err != nil {
		return nil, fmt.Errorf("%w:%w", ErrPrepareStmt, err)
	}

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, fmt.Errorf("%w:%w", ErrExecuteQuery, err)
	}
	defer rows.Close()

	var results []any
	for rows.Next() {
		item, err := scanner(rows)
		if err != nil {
			continue
		}
		results = append(results, item)
	}

	return results, nil
}
