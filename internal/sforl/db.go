package sforl

import (
	"context"
	"database/sql"
	"fmt"
	"log"
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
	selectMatchDrawStr   = fmt.Sprintf(`SELECT %[2]s, %[3]s, %[4]s, %[5]s, %[6]s, %[7]s, %[8]s, %[9]s, %[10]s, %[11]s, %[12]s FROM %[1]s WHERE %[4]s=? OR %[5]s=? OR %[6]s=? OR %[7]s=? OR %[8]s=? OR %[9]s=?`, tblName, drawDate, dayOfWeek, ball1, ball2, ball3, ball4, ball5, luckyBall, ballset, machine, drawNo)
)

func createTable(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, createTableStmtStr)
	if err != nil {
		return fmt.Errorf("%w-%s", dbutil.ErrDBCreateTbl, err.Error())
	}
	return nil
}

func persistsDrawChan(ctx context.Context, db *sql.DB, dc <-chan drawChan) error {
	stmt, err := prepareInsertDrawStmt(ctx, db)
	if err != nil {
		return err
	}
	for c := range dc {
		if c.Err != nil {
			log.Println(c.Err)
			continue
		}
		_, err = insertDraw(ctx, stmt, c.Draw)
		if err != nil {
			log.Println(err)
			continue
		}
	}
	return nil
}

func listAll(ctx context.Context, db *sql.DB) ([]Draw, error) {
	var draws []Draw
	row, err := db.QueryContext(ctx, selectAllDrawStmtStr)
	if err != nil {
		return nil, fmt.Errorf("%w-%s", dbutil.ErrDBQueryTbl, err.Error())
	}
	for row.Next() {
		d := Draw{}
		var drawDate int
		err := row.Scan(&drawDate, &d.DayOfWeek, &d.Ball1, &d.Ball2, &d.Ball3, &d.Ball4, &d.Ball5, &d.LifeBall, &d.BallSet, &d.Machine, &d.DrawNo)
		if err != nil {
			return nil, fmt.Errorf("%w-%s", dbutil.ErrDBQueryTbl, err.Error())
		}
		d.DrawDate = time.Unix(int64(drawDate), 0)
		draws = append(draws, d)
	}
	return draws, nil
}

func prepareInsertDrawStmt(ctx context.Context, db *sql.DB) (*sql.Stmt, error) {
	stmt, err := db.PrepareContext(ctx, insertDrawStmtStr)
	if err != nil {
		return nil, fmt.Errorf("%w-%s", dbutil.ErrDBPrepareStmt, err.Error())
	}
	return stmt, nil
}

func insertDraw(ctx context.Context, stmt *sql.Stmt, d Draw) (sql.Result, error) {
	result, err := stmt.ExecContext(ctx, d.DrawDate.Unix(), d.DayOfWeek, d.Ball1, d.Ball2, d.Ball3, d.Ball4, d.Ball5, d.LifeBall, d.BallSet, d.Machine, d.DrawNo)
	if err != nil {
		return nil, fmt.Errorf("%w-%s", dbutil.ErrDBInsertTbl, err.Error())
	}
	return result, nil
}

func matchDraw(ctx context.Context, db *sql.DB, ball1 uint, ball2 uint, ball3 uint, ball4 uint, ball5 uint, lifeball uint) ([]Draw, error) {
	var draws []Draw
	row, err := db.QueryContext(ctx, selectMatchDrawStr, ball1, ball2, ball3, ball4, ball5, lifeball)
	if err != nil {
		return nil, fmt.Errorf("%w-%s", dbutil.ErrDBQueryTbl, err.Error())
	}
	for row.Next() {
		d := Draw{}
		var drawDate int
		err := row.Scan(&drawDate, &d.DayOfWeek, &d.Ball1, &d.Ball2, &d.Ball3, &d.Ball4, &d.Ball5, &d.BallSet, &d.Machine, &d.DrawNo)
		if err != nil {
			return nil, fmt.Errorf("%w-%s", dbutil.ErrDBQueryTbl, err.Error())
		}
		d.DrawDate = time.Unix(int64(drawDate), 0)
		draws = append(draws, d)
	}
	return draws, nil
}
