package sforl

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"paulwizviz/lotterystat/internal/dbutil"
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

// SQLite

var (
	createSQLiteTableSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (%s INTEGER,%s INTEGER,%s INTEGER,%s INTEGER,%s INTEGER,%s INTEGER,%s INTEGER,%s INTEGER, %s TEXT,%s TEXT,%s INTEGER PRIMARY KEY)`, tblName, drawDate, dayOfWeek, ball1, ball2, ball3, ball4, ball5, luckyBall, ballset, machine, drawNo)
	//insertSQLiteDrawSQL  = fmt.Sprintf(`INSERT INTO %s (%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )`, tblName, drawDate, dayOfWeek, ball1, ball2, ball3, ball4, ball5, luckyBall, ballset, machine, drawNo)
	//selectSQLiteAllDrawSQL = fmt.Sprintf(`SELECT * FROM %s`, tblName)
)

func CreateSQLiteTable(ctx context.Context, db *sql.DB) error {
	return createSQLiteTable(ctx, db)
}

func createSQLiteTable(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, createSQLiteTableSQL)
	if err != nil {
		return fmt.Errorf("%w-%s", dbutil.ErrDBCreateTbl, err.Error())
	}
	return nil
}

// PSQL

var (
	createPSQLTableSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (%s INT,%s INT,%s INT,%s INT,%s INT,%s INTEGER,%s INT,%s INT, %s VARCHAR(64),%s VARCHAR(64),%s INT PRIMARY KEY)`, tblName, drawDate, dayOfWeek, ball1, ball2, ball3, ball4, ball5, luckyBall, ballset, machine, drawNo)
)

func createPSQLTable(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, createPSQLTableSQL)
	if err != nil {
		return fmt.Errorf("%w-%s", dbutil.ErrDBCreateTbl, err.Error())
	}
	return nil
}

func CreatePSQLTable(ctx context.Context, db *sql.DB) error {
	return createPSQLTable(ctx, db)
}

// Common for SQLite and PSQL

var (
	inserDrawSQL = fmt.Sprintf(`INSERT INTO %s (%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`, tblName, drawDate, dayOfWeek, ball1, ball2, ball3, ball4, ball5, luckyBall, ballset, machine, drawNo)
)

func prepInsertDrawStmt(ctx context.Context, db *sql.DB) (*sql.Stmt, error) {
	stmt, err := db.PrepareContext(ctx, inserDrawSQL)
	if err != nil {
		log.Println(inserDrawSQL)
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

func persistsDraw(ctx context.Context, db *sql.DB, dc <-chan DrawChan) error {
	stmt, err := prepInsertDrawStmt(ctx, db)
	if err != nil {
		return err
	}
	for c := range dc {
		if c.Err != nil {
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

func countBallSQL() string {
	return fmt.Sprintf("SELECT COUNT(*) FROM %[1]s WHERE %[2]s=$1 OR %[3]s=$1 OR %[4]s=$1 OR %[5]s=$1 OR %[6]s=$1;", tblName, ball1, ball2, ball3, ball4, ball5)
}

func prepCountBallStmt(ctx context.Context, db *sql.DB) (*sql.Stmt, error) {
	stmt, err := db.PrepareContext(ctx, countBallSQL())
	if err != nil {
		return nil, fmt.Errorf("%w-%s", dbutil.ErrDBPrepareStmt, err.Error())
	}
	return stmt, nil
}

func countBall(ctx context.Context, stmt *sql.Stmt, num uint8) (int, error) {
	rows, err := stmt.QueryContext(ctx, num)
	if err != nil {
		return 0, err
	}

	var count int
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			break
		}
	}
	return count, nil
}
