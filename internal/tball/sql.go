package tball

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/paulwizviz/lotterystat/internal/sqlops"
)

const (
	tblName   = "tball"
	drawDate  = "draw_date"
	dayOfWeek = "day_of_week"
	ball1     = "ball1"
	ball2     = "ball2"
	ball3     = "ball3"
	ball4     = "ball4"
	ball5     = "ball5"
	tball     = "tball"
	ballset   = "ball_set"
	machine   = "machine"
	drawNo    = "draw_no"
)

// SQL string statements
var (
	createTableSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
        %s INTEGER, %s INTEGER, %s INTEGER, %s INTEGER, %s INTEGER, %s INTEGER, %s INTEGER, %s INTEGER, %s INTEGER,%s TEXT,%s INTEGER PRIMARY KEY)`,
		tblName, drawDate, dayOfWeek, ball1, ball2, ball3, ball4, ball5, tball, ballset, machine, drawNo)

	writeDrawSQL = fmt.Sprintf(`INSERT INTO %s (
	    %s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`,
		tblName, drawDate, dayOfWeek, ball1, ball2, ball3, ball4, ball5, tball, ballset, machine, drawNo)

	selectAllDrawSQL = fmt.Sprintf(`SELECT * FROM %s`, tblName)

	countBallSQL = fmt.Sprintf(`SELECT COUNT(*) FROM %[1]s 
	    WHERE %[2]s=$1 OR %[3]s=$1 OR %[4]s=$1 OR %[5]s=$1 OR %[6]s=$1;`,
		tblName, ball1, ball2, ball3, ball4, ball5)

	countTBallSQL = fmt.Sprintf("SELECT COUNT(*) FROM %[1]s WHERE %[2]s=$1;", tblName, tball)
)

var CreateTableFn sqlops.TblCreatorFunc = func(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, createTableSQL)
	if err != nil {
		return err
	}
	return nil
}
