package euro

import (
	"context"
	"database/sql"
	"fmt"
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

var (
	createTableStmtStr      = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (%s INTEGER,%s INTEGER,%s INTEGER,%s INTEGER,%s INTEGER,%s INTEGER,%s INTEGER,%s INTEGER,%s INTEGER,%s TEXT,%s INTEGER PRIMARY KEY)`, tblName, drawDate, dayOfWeek, ball1, ball2, ball3, ball4, ball5, luckyStar1, luckyStar2, ukMarker, drawNo)
	insertDrawStmtStr       = fmt.Sprintf(`INSERT INTO %s (%s,%s,%s,%s,%s, %s,%s,%s,%s,%s,%s) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )`, tblName, drawDate, dayOfWeek, ball1, ball2, ball3, ball4, ball5, luckyStar1, luckyStar2, ukMarker, drawNo)
	selectAllStmtStr        = fmt.Sprintf(`SELECT * FROM %s`, tblName)
	selectMatchDrawStmtStr  = fmt.Sprintf(`SELECT %[2]s, %[3]s, %[4]s, %[5]s, %[6]s, %[7]s, %[8]s, %[9]s, %[10]s, %[11]s, %[12]s FROM %[1]s WHERE %[4]s=? OR %[5]s=? OR %[6]s=? OR %[7]s=? OR %[8]s=? OR %[9]s=? OR %[10]s=?`, tblName, drawDate, dayOfWeek, ball1, ball2, ball3, ball4, ball5, luckyStar1, luckyStar2, ukMarker, drawNo)
	selectCountBallsStmtStr = fmt.Sprintf(`SELECT COUNT(*) FROM %[1]s WHERE %[2]s=? OR %[3]s=? OR %[4]s=? OR %[5]s=? OR %[6]s=?`, tblName, ball1, ball2, ball3, ball4, ball5)
	selectCountStarsStmtStr = fmt.Sprintf(`SELECT COUNT(*) FROM %[1]s WHERE %[2]s=? OR %[3]s=?`, tblName, luckyStar1, luckyStar2)
)

func createTable(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, createTableStmtStr)
	if err != nil {
		return fmt.Errorf("%w-%s", dbutil.ErrDBCreateTbl, err.Error())
	}
	return nil
}

func persistsDrawChan(ctx context.Context, db *sql.DB, dc <-chan DrawChan) error {
	stmt, err := prepareInsertDrawStmt(ctx, db)
	if err != nil {
		return err
	}
	defer stmt.Close()
	for c := range dc {
		if c.Err != nil {
			continue
		}
		_, err = insertDraw(ctx, stmt, c.Draw)
		if err != nil {
			continue
		}
	}
	return nil
}

func listAllDraw(ctx context.Context, db *sql.DB) ([]Draw, error) {
	var draws []Draw
	row, err := db.QueryContext(ctx, selectAllStmtStr)
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

func prepareInsertDrawStmt(ctx context.Context, db *sql.DB) (*sql.Stmt, error) {
	stmt, err := db.PrepareContext(ctx, insertDrawStmtStr)
	if err != nil {
		return nil, fmt.Errorf("%w-%s", dbutil.ErrDBPrepareStmt, err.Error())
	}
	return stmt, nil
}

func insertDraw(ctx context.Context, stmt *sql.Stmt, d Draw) (sql.Result, error) {
	result, err := stmt.ExecContext(ctx, d.DrawDate.Unix(), d.DayOfWeek, d.Ball1, d.Ball2, d.Ball3, d.Ball4, d.Ball5, d.LS1, d.LS2, d.UKMarker, d.DrawNo)
	if err != nil {
		return nil, fmt.Errorf("%w-%s", dbutil.ErrDBInsertTbl, err.Error())
	}
	return result, nil
}

func matchBets(ctx context.Context, db *sql.DB, bets []Bet) ([]MatchedDraw, error) {
	stmt, err := prepareMatchDrawStmt(ctx, db)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var mds []MatchedDraw
	for _, b := range bets {
		b := b
		draws, err := matchDraw(ctx, stmt, b.Ball1, b.Ball2, b.Ball3, b.Ball4, b.Ball5, b.LS1, b.LS2)
		if err == nil {
			for _, d := range draws {
				d := d
				md := MatchedDraw{
					Bet:  b,
					Draw: d,
				}
				if d.Ball1 == b.Ball1 {
					md.Balls = append(md.Balls, d.Ball1)
				}
				if d.Ball2 == b.Ball2 {
					md.Balls = append(md.Balls, d.Ball2)
				}
				if d.Ball3 == b.Ball3 {
					md.Balls = append(md.Balls, d.Ball3)
				}
				if d.Ball4 == b.Ball4 {
					md.Balls = append(md.Balls, d.Ball4)
				}
				if d.Ball5 == b.Ball5 {
					md.Balls = append(md.Balls, d.Ball5)
				}
				if d.LS1 == b.LS1 {
					md.LuckyStars = append(md.LuckyStars, d.LS1)
				}
				if d.LS2 == b.LS2 {
					md.LuckyStars = append(md.LuckyStars, d.LS2)
				}
				mds = append(mds, md)
			}
		}
	}
	return mds, nil
}

func prepareMatchDrawStmt(ctx context.Context, db *sql.DB) (*sql.Stmt, error) {
	stmt, err := db.PrepareContext(ctx, selectMatchDrawStmtStr)
	if err != nil {
		return nil, err
	}
	return stmt, nil
}

func matchDraw(ctx context.Context, stmt *sql.Stmt, ball1 uint8, ball2 uint8, ball3 uint8, ball4 uint8, ball5 uint8, ls1 uint8, ls2 uint8) ([]Draw, error) {
	var draws []Draw
	row, err := stmt.QueryContext(ctx, ball1, ball2, ball3, ball4, ball5, ls1, ls2)
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

func prepareCountBallsStmt(ctx context.Context, db *sql.DB) (*sql.Stmt, error) {
	stmt, err := db.PrepareContext(ctx, selectCountBallsStmtStr)
	if err != nil {
		return nil, err
	}
	return stmt, nil
}

func countBall(ctx context.Context, stmt *sql.Stmt, ball uint8) (BallFreq, error) {
	row, err := stmt.QueryContext(ctx, ball, ball, ball, ball, ball)
	if err != nil {
		return BallFreq{}, nil
	}

	var count uint16
	for row.Next() {
		err = row.Scan(&count)
		if err != nil {
			return BallFreq{}, err
		}
	}
	return BallFreq{
		Ball:  ball,
		Count: uint16(count),
	}, nil
}

func prepareCountStarsStmt(ctx context.Context, db *sql.DB) (*sql.Stmt, error) {
	stmt, err := db.PrepareContext(ctx, selectCountStarsStmtStr)
	if err != nil {
		return nil, err
	}
	return stmt, nil
}

func countStars(ctx context.Context, stmt *sql.Stmt, star uint8) (StarFreq, error) {
	row, err := stmt.QueryContext(ctx, star, star)
	if err != nil {
		return StarFreq{}, nil
	}

	var count uint16
	for row.Next() {
		err = row.Scan(&count)
		if err != nil {
			return StarFreq{}, err
		}
	}
	return StarFreq{
		Star:  star,
		Count: uint16(count),
	}, nil
}
