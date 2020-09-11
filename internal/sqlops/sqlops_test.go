package sqlops_test

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/paulwizviz/lotterystat/internal/sqlops"
)

var createTableFn sqlops.TblCreatorFunc = func(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS draw(id PRIMARY KEY, ball1 INTEGER)")
	if err != nil {
		return err
	}
	return nil
}

func Example() {

	db, err := sqlops.NewSQLiteMem()
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx := context.TODO()

	err = sqlops.CreateTables(ctx, db, createTableFn)
	if err != nil {
		fmt.Println(err)
		return
	}

}
