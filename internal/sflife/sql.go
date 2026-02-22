package sflife

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/paulwizviz/lotterystat/internal/sqlops"
)

const (
	tblName   = "sflife"
	drawDate  = "draw_date"
	dayOfWeek = "day_of_week"
	ball1     = "ball1"
	ball2     = "ball2"
	ball3     = "ball3"
	ball4     = "ball4"
	ball5     = "ball5"
	lball     = "lball"
	ballset   = "ball_set"
	machine   = "machine"
	drawNo    = "draw_no"
)

var (
	createTableSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
        %s INTEGER, %s INTEGER, %s INTEGER, %s INTEGER, %s INTEGER, %s INTEGER, %s INTEGER, %s INTEGER,%s TEXT,%s TEXT,%s INTEGER PRIMARY KEY)`,
		tblName, drawDate, dayOfWeek, ball1, ball2, ball3, ball4, ball5, lball, ballset, machine, drawNo)

	CreateTableFn sqlops.TblCreator = func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, createTableSQL)
		if err != nil {
			return err
		}
		return nil
	}
)

var (
	writeDrawSQL = fmt.Sprintf(`INSERT INTO %s (
	    %s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`,
		tblName, drawDate, dayOfWeek, ball1, ball2, ball3, ball4, ball5, lball, ballset, machine, drawNo)

	writeDrawRowFn = func(ctx context.Context, stmt *sql.Stmt, data any) error {
		d, ok := data.(Draw)
		if !ok {
			return fmt.Errorf("%w: invalid argument type", sqlops.ErrExecuteWriter)
		}
		_, err := stmt.ExecContext(ctx, d.DrawDate, d.DayOfWeek, d.Ball1, d.Ball2, d.Ball3, d.Ball4, d.Ball5, d.LBall, d.BallSet, d.Machine, d.DrawNo)
		if err != nil {
			return fmt.Errorf("%w:%w", sqlops.ErrExecuteWriter, err)
		}
		return nil
	}
)

func PersistsDraw(ctx context.Context, db *sql.DB, data Draw) error {
	return sqlops.Writer(ctx, db, writeDrawSQL, []any{data}, writeDrawRowFn)
}

var (
	selectAllDrawSQL = fmt.Sprintf(`SELECT * FROM %s`, tblName)
)

func ListAllDraws(ctx context.Context, db *sql.DB) ([]Draw, error) {

	result, err := sqlops.Query(ctx, db, func(rows *sql.Rows) (any, error) {
		d := Draw{}
		var drawDate string
		err := rows.Scan(&drawDate, &d.DayOfWeek, &d.Ball1, &d.Ball2, &d.Ball3, &d.Ball4, &d.Ball5, &d.LBall, &d.BallSet, &d.Machine, &d.DrawNo)
		if err != nil {
			return nil, fmt.Errorf("%w:%w", sqlops.ErrExecuteQuery, err)
		}
		d.DrawDate, err = time.Parse("2006-01-02 15:04:05 -0700 MST", drawDate)
		if err != nil {
			return nil, err
		}
		return d, nil
	}, selectAllDrawSQL)
	if err != nil {
		return nil, err
	}

	draws := []Draw{}
	for _, item := range result {
		d := item.(Draw)
		draws = append(draws, d)
	}

	return draws, nil
}

var (
	countBallSQL = fmt.Sprintf(`SELECT COUNT(*) FROM %[1]s 
	    WHERE %[2]s=$1 OR %[3]s=$1 OR %[4]s=$1 OR %[5]s=$1 OR %[6]s=$1;`,
		tblName, ball1, ball2, ball3, ball4, ball5)
)

type BallFrequency struct {
	Ball      uint
	Frequency uint
}

func CalculateBallFreq(ctx context.Context, db *sql.DB) ([]BallFrequency, error) {

	ballFreqs := []BallFrequency{}
	ball := 0
	for range 47 {
		ball = ball + 1

		bf := BallFrequency{
			Ball: uint(ball),
		}

		result, err := sqlops.Query(ctx, db, func(r *sql.Rows) (any, error) {
			var count int
			if err := r.Scan(&count); err != nil {
				return nil, fmt.Errorf("%w: %v", sqlops.ErrExecuteQuery, err)
			}
			return count, nil
		}, countBallSQL, bf.Ball)

		if err != nil {
			return nil, err
		}

		bf.Frequency = uint(result[0].(int))
		ballFreqs = append(ballFreqs, bf)
	}

	return ballFreqs, nil
}

var (
	countLBallSQL = fmt.Sprintf("SELECT COUNT(*) FROM %[1]s WHERE %[2]s=$1;", tblName, lball)
)

type LBallFrequency struct {
	LBall     uint
	Frequency uint
}

func CalculateLBallFreq(ctx context.Context, db *sql.DB) ([]LBallFrequency, error) {
	lballFreqs := []LBallFrequency{}
	lBall := 0
	for range 10 {
		lBall = lBall + 1
		lballFreq := LBallFrequency{
			LBall: uint(lBall),
		}
		result, err := sqlops.Query(ctx, db, func(r *sql.Rows) (any, error) {
			var count int
			if err := r.Scan(&count); err != nil {
				return nil, fmt.Errorf("%w: %v", sqlops.ErrExecuteQuery, err)
			}
			return count, nil
		}, countLBallSQL, lballFreq.LBall)
		if err != nil {
			return nil, err
		}
		lballFreq.Frequency = uint(result[0].(int))
		lballFreqs = append(lballFreqs, lballFreq)
	}
	return lballFreqs, nil
}
