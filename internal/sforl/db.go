package sforl

import (
	"context"
	"database/sql"
	"fmt"
	"paulwizviz/lotterystat/internal/dbutil"
	"time"
)

const (
	tblName   = "set_for_life"
	drawDate  = "draw_date"
	dayOfWeek = "day_of_week"
	ball1     = "ball1"
	ball2     = "ball2"
	ball3     = "ball3"
	ball4     = "ball4"
	ball5     = "ball5"
	luckyBall = "lb"
	ballset   = "ball_set"
	machine   = "machine"
	drawNo    = "draw_no"
)

var (
	createTableStmtStr   = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (%s INTEGER,%s INTEGER,%s INTEGER,%s INTEGER,%s INTEGER,%s INTEGER,%s INTEGER,%s INTEGER, %s TEXT,%s TEXT,%s INTEGER PRIMARY KEY)`, tblName, drawDate, dayOfWeek, ball1, ball2, ball3, ball4, ball5, luckyBall, ballset, machine, drawNo)
	insertDrawStmtStr    = fmt.Sprintf(`INSERT INTO %s (%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )`, tblName, drawDate, dayOfWeek, ball1, ball2, ball3, ball4, ball5, luckyBall, ballset, machine, drawNo)
	selectAllDrawStmtStr = fmt.Sprintf(`SELECT * FROM %s`, tblName)
)

func CreateTable(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, createTableStmtStr)
	if err != nil {
		return fmt.Errorf("%w-%s", dbutil.ErrDBCreateTbl, err.Error())
	}
	return nil
}

func PrepareInsertDrawStmt(ctx context.Context, db *sql.DB) (*sql.Stmt, error) {
	return db.PrepareContext(ctx, insertDrawStmtStr)
}

func InsertDraw(ctx context.Context, stmt *sql.Stmt, d Draw) (sql.Result, error) {
	result, err := stmt.ExecContext(ctx, d.DrawDate.Unix(), d.DayOfWeek, d.Ball1, d.Ball2, d.Ball3, d.Ball4, d.Ball5, d.LifeBall, d.Machine, d.DrawNo)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func ListAll(ctx context.Context, db *sql.DB) ([]Draw, error) {
	var draws []Draw
	row, err := db.QueryContext(ctx, selectAllDrawStmtStr)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		d := Draw{}
		var drawDate int
		err := row.Scan(&drawDate, &d.DayOfWeek, &d.Ball1, &d.Ball2, &d.Ball3, &d.Ball4, &d.Ball5, &d.LifeBall, &d.Machine, &d.DrawNo)
		if err == nil {
			d.DrawDate = time.Unix(int64(drawDate), 0)
			draws = append(draws, d)
		}
	}
	return draws, nil
}
