package euro

import "time"

const (
	CSVUrl = "https://www.national-lottery.co.uk/results/euromillions/draw-history/csv"
)

// Draw represents a line from euro draw results
type Draw struct {
	DrawDate  time.Time    `json:"draw_date"`
	DayOfWeek time.Weekday `json:"day_of_week"`
	Ball1     uint8        `json:"ball1"`
	Ball2     uint8        `json:"ball2"`
	Ball3     uint8        `json:"ball3"`
	Ball4     uint8        `json:"ball4"`
	Ball5     uint8        `json:"ball5"`
	LS1       uint8        `json:"ls1"`
	LS2       uint8        `json:"ls2"`
	UKMarker  string       `json:"uk_marker"`
	DrawNo    uint64       `json:"draw_no"`
}

var (
	SQLiteCreateTblStr = `CREATE TABLE IF NOT EXISTS euro (
		draw_date INTEGER, 
		day_of_week INTEGER, 
		ball1 INTEGER, 
		ball2 INTEGER, 
		ball3 INTEGER, 
		ball4 INTEGER, 
		ball5 INTEGER, 
		ls1 INTEGER, 
		ls2 INTEGER, 
		uk_marker TEXT, 
		draw_no INTEGER PRIMARY KEY)`
	SQLiteInsertEuroStr = `INSERT INTO euro (
		draw_date, 
		day_of_week, 
		ball1, 
		ball2, 
		ball3, 
		ball4, 
		ball5, 
		ls1, 
		ls2, 
		uk_marker, 
		draw_no) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )`
	SQLiteListAllEuroStr = `SELECT * FROM euro`
)
