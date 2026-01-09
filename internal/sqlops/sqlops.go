package sqlops

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDBConn, err)
	}
	return db, nil
}

// NewSQLiteFile instantiate a file based SQLite
func NewSQLiteFile(f string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", f)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDBConn, err)
	}
	return db, nil
}

// TblCreator is an interface type to support
// the creation of database table
type TblCreator interface {
	Create(context.Context, *sql.Tx) error
}

// TblCreatorFunc is a function type to help create
// db table
type TblCreatorFunc func(context.Context, *sql.Tx) error

func (t TblCreatorFunc) Create(ctx context.Context, tx *sql.Tx) error {
	return t(ctx, tx)
}

// CreateTables is a wrapper function to create Table based on the
// a slice of table creators
func CreateTables(ctx context.Context, db *sql.DB, creators ...TblCreator) error {
	tx, err := db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return fmt.Errorf("%w: %v", ErrCreateTxn, err)
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		} else if err != nil {
			tx.Rollback()
		}
	}()

	for _, creator := range creators {
		err := creator.Create(ctx, tx)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrCreateTbl, err)
		}
	}

	tx.Commit()
	return nil
}

// StmtBuilder is a builder to support the creation of prepared SQL statement
type StmtBuilder struct {
	stmt *sql.Stmt
}

func (s *StmtBuilder) PrepareStatement(ctx context.Context, db *sql.DB, query string) error {
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrPrepareStmt, err)
	}
	s.stmt = stmt
	return nil
}

func (s *StmtBuilder) Close() error {
	if s.stmt == nil {
		return nil // Already closed or never prepared, nothing to do
	}

	err := s.stmt.Close()
	s.stmt = nil // Clear the statement after closing it
	if err != nil {
		return fmt.Errorf("%w: %v", ErrCloseStmt, err)
	}
	return nil
}

type TableWriter struct {
	*StmtBuilder
}

func (t *TableWriter) Execute(ctx context.Context, args ...any) error {
	_, err := t.StmtBuilder.stmt.ExecContext(ctx, args)
	if err != nil {
		return err
	}
	return nil
}
