package draw

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"reflect"
	"strings"
	"time"
)

var (
	ErrDBCreateTbl = errors.New("unable to create table")
)

// TblStmtCreator is an interface for the implementation
// of function to generate create table statements
type TblStmtCreator[T DrawType] interface {
	Create(typ *T) string
}

// CreateTblStmtfuncL is a logging middleware for
// func CreateTblStmt
type CreateTblStmtfuncL[T DrawType] func(typ *T) string

func (c CreateTblStmtfuncL[T]) Create(typ *T) string {
	stmt := c(typ)
	slog.Debug("TblStmtCreator", slog.String("Statement", stmt))
	return stmt
}

// CreateTblStmt is an implementation to generate create
// table statement
func CreateTblStmt[T DrawType](v *T) string {

	t := reflect.TypeOf(v)
	tblName := strings.Split(fmt.Sprintf("%v", t), ".")[1]
	tags := sqliteTags(v)

	stmt := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (", tblName)
	for _, tag := range tags {
		tagElemes := strings.Split(tag.Tag, ",")
		stmt = fmt.Sprintf("%s %s %s,", stmt, tagElemes[0], tagElemes[1])
	}
	stmt = stmt[0 : len(stmt)-1] // remove last comma
	stmt = fmt.Sprintf(`%s PRIMARY KEY`, stmt)
	stmt = fmt.Sprintf(`%s )`, stmt)

	return stmt
}

// TblCreator is an interface representing operations to
// create draw table
type TblCreator[T DrawType] interface {
	Create(db *sql.DB, creator TblStmtCreator[T], typ *T) (sql.Result, error)
}

// CreateTblFuncL is a logging middleware for func CreateTbl
type CreateTblFuncL[T DrawType] func(db *sql.DB, creator TblStmtCreator[T], typ *T) (sql.Result, error)

func (c CreateTblFuncL[T]) Create(db *sql.DB, creator TblStmtCreator[T], typ *T) (sql.Result, error) {
	logMsg := "Create Table"
	result, err := c(db, creator, typ)
	if err != nil {
		slog.Info(logMsg, slog.String("create table error", err.Error()))
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		slog.Info(logMsg, slog.String("last insert ID error", err.Error()))
	}

	row, err := result.RowsAffected()
	if err != nil {
		slog.Info(logMsg, slog.String("last insert ID error", err.Error()))
	}

	slog.Info(logMsg,
		slog.Int64("last insert ID", lastID),
		slog.Int64("row affected ID", row),
	)
	return result, err
}

// CreateTbl is an implementation of operations to create Table
func CreateTbl[T DrawType](db *sql.DB, creator TblStmtCreator[T], typ *T) (sql.Result, error) {
	stmt := creator.Create(typ)
	result, err := db.Exec(stmt)
	if err != nil {
		return result, fmt.Errorf("%w: %s", ErrDBCreateTbl, err.Error())
	}
	return result, nil
}

type structTag struct {
	FieldName string
	Tag       string
}

func sqliteTags[T DrawType](typ *T) []structTag {
	ev := reflect.Indirect(reflect.ValueOf(typ))
	tags := []structTag{}
	for i := 0; i < ev.Type().NumField(); i++ {
		tag := structTag{}
		tag.FieldName = ev.Type().Field(i).Name
		t := ev.Type().Field(i).Tag
		tElems := strings.Split(string(t), " ")
		for _, tElem := range tElems {
			if strings.Contains(tElem, "sqlite") {
				sElems := strings.Split(tElem, ":")
				tag.Tag = sElems[1][1 : len(sElems[1])-1]
			}
		}
		tags = append(tags, tag)
	}
	return tags
}

func createInsertStmt[T DrawType](val *T) string {
	t := reflect.TypeOf(val)
	tblName := strings.Split(fmt.Sprintf("%v", t), ".")[1]
	tags := sqliteTags(val)
	v := reflect.Indirect(reflect.ValueOf(val))

	stmt := fmt.Sprintf("INSERT INTO %s (", tblName)
	for _, tag := range tags {
		t1 := strings.Split(tag.Tag, ",")
		stmt = fmt.Sprintf("%s %s,", stmt, t1[0])
	}
	stmt = stmt[0 : len(stmt)-1]
	stmt = fmt.Sprintf("%s) VALUES (", stmt)
	for _, tag := range tags {
		fval := v.FieldByName(tag.FieldName)
		switch fv := fval.Interface().(type) {
		case time.Time:
			ut := fv.Unix()
			stmt = fmt.Sprintf("%s %d,", stmt, ut)
		case time.Weekday:
			stmt = fmt.Sprintf("%s %d,", stmt, uint64(fv))
		case uint8, uint64:
			stmt = fmt.Sprintf("%s %d,", stmt, fv)
		case string:
			stmt = fmt.Sprintf(`%s "%s",`, stmt, fval)
		default:
			stmt = fmt.Sprintf(`%v "%s",`, stmt, fval)
		}
	}
	stmt = stmt[0 : len(stmt)-1]
	stmt = fmt.Sprintf("%s)", stmt)

	return stmt
}
