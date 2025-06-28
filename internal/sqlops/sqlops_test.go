package sqlops_test

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"

	"github.com/paulwizviz/lotterystat/internal/sqlops"
)

var (
	sqlitePerson = `CREATE TABLE IF NOT EXISTS person (
first_name TEXT,
surname TEXT,
age INTEGER)`

	pgPerson = `CREATE TABLE IF NOT EXISTS person (
first_name VARCHAR(10),
surname VARCHAR(10),
age INT)`

	createPersonTblFunc sqlops.TblCreatorFunc = func(ctx context.Context, db *sql.DB) error {
		dbType := sqlops.DriverType(db)
		switch dbType {
		case sqlops.SQLiteType:
			_, err := db.ExecContext(ctx, pgPerson)
			if err != nil {
				return fmt.Errorf("%w-person table", sqlops.ErrCreateTbl)
			}
		case sqlops.PSQLType:
			_, err := db.ExecContext(ctx, pgPerson)
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("%w-driver %s unsupported", sqlops.ErrCreateTbl, reflect.TypeOf(db).String())
		}
		return nil
	}
)

func Example() {
	db, err := sqlops.NewSQLiteMem()
	if err != nil {
		return
	}
	defer db.Close()
	err = sqlops.CreateTable(context.TODO(), db, createPersonTblFunc)
	fmt.Println(err)

	// Output:
}
