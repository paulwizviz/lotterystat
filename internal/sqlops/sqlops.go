package sqlops

import (
	"context"
	"database/sql"
	"fmt"
)

// TblCreator is an interface type to support
// the creation of database table
type TblCreator interface {
	CreateTable(context.Context, *sql.DB) error
}

// TblCreatorFunc is a function type to help create
// db table
type TblCreatorFunc func(context.Context, *sql.DB) error

func (t TblCreatorFunc) CreateTable(ctx context.Context, db *sql.DB) error {
	return t(ctx, db)
}

// Create table is a wrapper function to create Table based on the
// table creator interface
func CreateTable(ctx context.Context, db *sql.DB, creator TblCreator) error {
	err := creator.CreateTable(ctx, db)
	if err != nil {
		return fmt.Errorf("%w-%v", ErrCreateTbl, err)
	}
	return nil
}
