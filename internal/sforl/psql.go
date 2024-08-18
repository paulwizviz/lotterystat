package sforl

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"paulwizviz/lotterystat/internal/dbutil"
)

var (
	createPSQLTableSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (%s INT,%s INT,%s INT,%s INT,%s INT,%s INTEGER,%s INT,%s INT, %s VARCHAR(64),%s VARCHAR(64),%s INT PRIMARY KEY)`, tblName, drawDate, dayOfWeek, ball1, ball2, ball3, ball4, ball5, luckyBall, ballset, machine, drawNo)
	insertPSQLDrawSQL  = fmt.Sprintf(`INSERT INTO %s (%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`, tblName, drawDate, dayOfWeek, ball1, ball2, ball3, ball4, ball5, luckyBall, ballset, machine, drawNo)
)

func CreatePSQLTable(ctx context.Context, db *sql.DB) error {
	return createPSQLTable(ctx, db)
}

func createPSQLTable(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, createPSQLTableSQL)
	if err != nil {
		return fmt.Errorf("%w-%s", dbutil.ErrDBCreateTbl, err.Error())
	}
	return nil
}

func PersistsPSQLCSV(ctx context.Context, db *sql.DB, nworkers int) error {
	return persistsPSQLCSV(ctx, db, nworkers)
}

func persistsPSQLDraw(ctx context.Context, db *sql.DB, dc <-chan DrawChan) error {
	stmt, err := prepPSQLInsertDrawStmt(ctx, db)
	if err != nil {
		return err
	}
	for c := range dc {
		if c.Err != nil {
			continue
		}
		_, err = insertPSQLDraw(ctx, stmt, c.Draw)
		if err != nil {
			log.Println(err)
			continue
		}
	}
	return nil
}

func prepPSQLInsertDrawStmt(ctx context.Context, db *sql.DB) (*sql.Stmt, error) {
	stmt, err := db.PrepareContext(ctx, insertPSQLDrawSQL)
	if err != nil {
		log.Println(insertPSQLDrawSQL)
		return nil, fmt.Errorf("%w-%s", dbutil.ErrDBPrepareStmt, err.Error())
	}
	return stmt, nil
}

func insertPSQLDraw(ctx context.Context, stmt *sql.Stmt, d Draw) (sql.Result, error) {
	result, err := stmt.ExecContext(ctx, d.DrawDate.Unix(), d.DayOfWeek, d.Ball1, d.Ball2, d.Ball3, d.Ball4, d.Ball5, d.LifeBall, d.BallSet, d.Machine, d.DrawNo)
	if err != nil {
		return nil, fmt.Errorf("%w-%s", dbutil.ErrDBInsertTbl, err.Error())
	}
	return result, nil
}
