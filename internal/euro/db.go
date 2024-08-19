package euro

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"paulwizviz/lotterystat/internal/dbutil"
	"time"
)

const (
	tblName    = "euro"
	drawDate   = "draw_date"
	dayOfWeek  = "day_of_week"
	ball1      = "ball1"
	ball2      = "ball2"
	ball3      = "ball3"
	ball4      = "ball4"
	ball5      = "ball5"
	luckyStar1 = "ls1"
	luckyStar2 = "ls2"
	ukMarker   = "uk_marker"
	drawNo     = "draw_no"
)

// SQLite

var (
	createTableSQLiteSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (%s INTEGER,%s INTEGER,%s INTEGER,%s INTEGER,%s INTEGER,%s INTEGER,%s INTEGER,%s INTEGER,%s INTEGER,%s TEXT,%s INTEGER PRIMARY KEY)`, tblName, drawDate, dayOfWeek, ball1, ball2, ball3, ball4, ball5, luckyStar1, luckyStar2, ukMarker, drawNo)
	insertDrawSQLiteSQL  = fmt.Sprintf(`INSERT INTO %s (%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )`, tblName, drawDate, dayOfWeek, ball1, ball2, ball3, ball4, ball5, luckyStar1, luckyStar2, ukMarker, drawNo)
	selectAllSQLiteSQL   = fmt.Sprintf(`SELECT * FROM %s`, tblName)
)

func CreateSQLiteTable(ctx context.Context, db *sql.DB) error {
	return createSQLiteTable(ctx, db)
}

func createSQLiteTable(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, createTableSQLiteSQL)
	if err != nil {
		return fmt.Errorf("%w-%s", dbutil.ErrDBCreateTbl, err.Error())
	}
	return nil
}

func persistsSQLiteDrawChan(ctx context.Context, db *sql.DB, dc <-chan DrawChan) error {
	stmt, err := prepSQLiteInsertDrawStmt(ctx, db)
	if err != nil {
		return err
	}
	defer stmt.Close()
	for c := range dc {
		if c.Err != nil {
			continue
		}
		_, err = insertSQLiteDraw(ctx, stmt, c.Draw)
		if err != nil {
			log.Println(err)
			continue
		}
	}
	return nil
}

func listSQLiteAllDraw(ctx context.Context, db *sql.DB) ([]Draw, error) {
	var draws []Draw
	row, err := db.QueryContext(ctx, selectAllSQLiteSQL)
	if err != nil {
		return nil, fmt.Errorf("%w-%s", dbutil.ErrDBQueryTbl, err.Error())
	}
	for row.Next() {
		d := Draw{}
		var drawDate int
		err := row.Scan(&drawDate, &d.DayOfWeek, &d.Ball1, &d.Ball2, &d.Ball3, &d.Ball4, &d.Ball5, &d.LS1, &d.LS2, &d.UKMarker, &d.DrawNo)
		if err != nil {
			return nil, fmt.Errorf("%w-%s", dbutil.ErrDBQueryTbl, err.Error())
		}
		d.DrawDate = time.Unix(int64(drawDate), 0)
		draws = append(draws, d)
	}
	return draws, nil
}

func prepSQLiteInsertDrawStmt(ctx context.Context, db *sql.DB) (*sql.Stmt, error) {
	stmt, err := db.PrepareContext(ctx, insertDrawSQLiteSQL)
	if err != nil {
		return nil, fmt.Errorf("%w-%s", dbutil.ErrDBPrepareStmt, err.Error())
	}
	return stmt, nil
}

func insertSQLiteDraw(ctx context.Context, stmt *sql.Stmt, d Draw) (sql.Result, error) {
	result, err := stmt.ExecContext(ctx, d.DrawDate.Unix(), d.DayOfWeek, d.Ball1, d.Ball2, d.Ball3, d.Ball4, d.Ball5, d.LS1, d.LS2, d.UKMarker, d.DrawNo)
	if err != nil {
		return nil, fmt.Errorf("%w-%s", dbutil.ErrDBInsertTbl, err.Error())
	}
	return result, nil
}

// PSQL

var (
	createPSQLTableSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (%s INT,%s INT,%s INT,%s INT,%s INT,%s INT,%s INT,%s INT,%s INT,%s VARCHAR(256),%s INT PRIMARY KEY)`, tblName, drawDate, dayOfWeek, ball1, ball2, ball3, ball4, ball5, luckyStar1, luckyStar2, ukMarker, drawNo)
	insertPSQLDrawSQL  = fmt.Sprintf(`INSERT INTO %s (%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`, tblName, drawDate, dayOfWeek, ball1, ball2, ball3, ball4, ball5, luckyStar1, luckyStar2, ukMarker, drawNo)
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
	result, err := stmt.ExecContext(ctx, d.DrawDate.Unix(), d.DayOfWeek, d.Ball1, d.Ball2, d.Ball3, d.Ball4, d.Ball5, d.LS1, d.LS2, d.UKMarker, d.DrawNo)
	if err != nil {
		return nil, fmt.Errorf("%w-%s", dbutil.ErrDBInsertTbl, err.Error())
	}
	return result, nil
}
