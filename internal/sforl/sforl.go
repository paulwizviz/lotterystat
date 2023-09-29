package sforl

import "time"

// Draw represents a draw from Set for Life
type Draw struct {
	DrawDate  time.Time    `json:"draw_date"`
	DayOfWeek time.Weekday `json:"day_of_week"`
	Ball1     uint8        `json:"ball1"`
	Ball2     uint8        `json:"ball2"`
	Ball3     uint8        `json:"ball3"`
	Ball4     uint8        `json:"ball4"`
	Ball5     uint8        `json:"ball5"`
	LifeBall  uint8        `json:"life_ball"`
	BallSet   string       `json:"ball_set"`
	Machine   string       `json:"machine"`
	DrawNo    uint64       `json:"draw_no"`
}

var (
	SQLiteCreateTblStr = `CREATE TABLE IF NOT EXISTS set_for_life (
		draw_date INTEGER, 
		day_of_week INTEGER, 
		ball1 INTEGER, 
		ball3 INTEGER, 
		ball4 INTEGER, 
		ball5 INTEGER, 
		lb INTEGER,  
		ball_set TEXT,
		machine TEXT,
		draw_no INTEGER PRIMARY KEY)`
	SQLiteInsertEuroStr = `INSERT INTO set_for_life (
		draw_date, 
		day_of_week, 
		ball1,
		ball2, 
		ball3,
		ball4,
		ball5,
		lb,
		ball_set,
		machine, 
		draw_no) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )`
)
