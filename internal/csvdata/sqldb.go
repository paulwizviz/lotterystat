package csvdata

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"
)

// TblCreator is an interface that wrap an operation to create Table
type TblCreator[T DrawType] interface {
	Create(*sql.DB, string, T) error
}

// TblCreateHdlFunc register a function to create an SQL Table to persists
// CSV data types
type TblCreateHdlFunc[T DrawType] func(*sql.DB, string, T) error

func (t TblCreateHdlFunc[T]) Create(db *sql.DB, tblname string, data T) error {
	return t(db, tblname, data)
}

// TblCreateHdl is the middleware responsible for the execution of a handler function
func TblCreateHdl[T DrawType](db *sql.DB, data T, hdl TblCreator[T]) error {

	tblname := reflect.TypeOf(data).Name()

	err := hdl.Create(db, tblname, data)
	if err != nil {
		return err
	}
	return nil
}

// TblWriter is an interface that wrap an operation to write CSV Data Types
// to an SQL Table
type TblWriter[T DrawType] interface {
	Write(*sql.DB, string, T) error
}

// TblWriterHdlFunc register a function to store CSV Data types to SQL Table
type TblWriterHdlFunc[T DrawType] func(db *sql.DB, tblName string, data T) error

func (t TblWriterHdlFunc[T]) Write(db *sql.DB, tblname string, data T) error {
	return t(db, tblname, data)
}

// TblWriterHdl is responsible for the execution of operation to store
// CSV data type in SQL Table
func TblWriterHdl[T DrawType](db *sql.DB, data T, hdl TblWriterHdlFunc[T]) error {

	tblname := reflect.TypeOf(data).Name()
	err := hdl.Write(db, tblname, data)
	if err != nil {
		return err
	}
	return nil
}

// TblReader is an interface that wraps an operation to read CSV Data Types
// from an SQL Table
type TblReader[T DrawType] interface {
	Read(*sql.DB) ([]T, error)
}

// TblReaderHdlFunc register a function to read CSV Data from SQL Table
type TblReaderHdlFunc[T DrawType] func(db *sql.DB) ([]T, error)

func (t TblReaderHdlFunc[T]) Read(db *sql.DB) ([]T, error) {
	return t(db)
}

func TblReadHdl[T DrawType](db *sql.DB, hdl TblReaderHdlFunc[T]) ([]T, error) {
	return hdl.Read(db)
}

var (

	// SQLiteCreateEuroDrawTbl is a default handler function to create a table in SQLite.
	//
	// Example statement
	// CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)
	SQLiteCreateEuroDrawTbl TblCreateHdlFunc[EuroDraw] = func(db *sql.DB, tblName string, data EuroDraw) error {
		tags := data.SQLiteTags()
		return createTbl(db, tblName, tags)
	}

	// SQLiteWriteEuroDrawTbl is a default operation to insert object into sqlite table
	//
	// Example statement
	// INSERT INTO people (firstname, lastname) VALUES (?, ?)
	SQLiteWriteEuroDrawTbl TblWriterHdlFunc[EuroDraw] = func(db *sql.DB, tblName string, data EuroDraw) error {
		tags := data.SQLiteTags()
		elem := reflect.ValueOf(&data).Elem()
		return insertIntoTbl(db, tblName, elem, tags)
	}

	// SQLiteReadEuroDrawTbl is a default operations to return all Euro Draws from database
	SQLiteReadEuroDrawTbl TblReaderHdlFunc[EuroDraw] = func(db *sql.DB) ([]EuroDraw, error) {
		tblname := reflect.TypeOf(EuroDraw{}).Name()
		stmtStr := fmt.Sprintf("SELECT * FROM %s", tblname)
		rows, err := db.Query(stmtStr)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		draws := []EuroDraw{}
		for rows.Next() {
			d := EuroDraw{}
			var dt int64
			err := rows.Scan(&dt, &d.DayOfWeek, &d.Ball1, &d.Ball2, &d.Ball3, &d.Ball4, &d.Ball5, &d.LS1, &d.LS2, &d.UKMarker, &d.EuroMarker, &d.DrawNo)
			if err != nil {
				break
			}
			d.DrawDate = time.Unix(dt, 0)
			draws = append(draws, d)
		}
		return draws, nil
	}
)

func createTbl(db *sql.DB, tblName string, tags []StructTag) error {
	stmtStr := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (", tblName)
	for _, tag := range tags {
		tagElemes := strings.Split(tag.Tag, ",")
		stmtStr = fmt.Sprintf("%s %s %s,", stmtStr, tagElemes[0], tagElemes[1])
	}
	stmtStr = stmtStr[0 : len(stmtStr)-1]
	stmtStr = fmt.Sprintf(`%s PRIMARY KEY`, stmtStr)
	stmtStr = fmt.Sprintf(`%s )`, stmtStr)
	stmt, err := db.Prepare(stmtStr)
	if err != nil {
		return fmt.Errorf("sqlite error %w-%w", err, ErrUnableToCreateTable)
	}
	defer stmt.Close()
	stmt.Exec()
	return nil
}

func insertIntoTbl(db *sql.DB, tblName string, elem reflect.Value, tags []StructTag) error {
	stmtStr := fmt.Sprintf("INSERT INTO %s (", tblName)
	for _, tag := range tags {
		t1 := strings.Split(tag.Tag, ",")
		stmtStr = fmt.Sprintf("%s %s,", stmtStr, t1[0])
	}
	stmtStr = stmtStr[0 : len(stmtStr)-1]
	stmtStr = fmt.Sprintf("%s) VALUES (", stmtStr)
	for _, tag := range tags {
		fval := elem.FieldByName(tag.FieldName)
		switch fv := fval.Interface().(type) {
		case time.Time:
			ut := fv.Unix()
			stmtStr = fmt.Sprintf("%s %d,", stmtStr, ut)
		case time.Weekday:
			stmtStr = fmt.Sprintf("%s %d,", stmtStr, uint64(fv))
		case uint8, uint64:
			stmtStr = fmt.Sprintf("%s %d,", stmtStr, fv)
		case string:
			stmtStr = fmt.Sprintf(`%s "%s",`, stmtStr, fval)
		default:
			stmtStr = fmt.Sprintf(`%v "%s",`, stmtStr, fval)
		}
	}
	stmtStr = stmtStr[0 : len(stmtStr)-1]
	stmtStr = fmt.Sprintf("%s)", stmtStr)
	stmt, err := db.Prepare(stmtStr)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("duplicate error - %s %w", err.Error(), ErrDuplicateEntry)
	}

	return nil
}
