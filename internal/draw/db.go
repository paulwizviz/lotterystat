package draw

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
)

var (
	ErrDBCreateTbl = errors.New("unable to create table")
)

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

type CreateTblStmtFunc[T DrawType] func(*T) string

// CreateTblStmt is an implementation to generate create
// table statement
func CreateTblStmt[T DrawType](v *T) string {
	t := reflect.TypeOf(v)
	tblName := strings.Split(fmt.Sprintf("%v", t), ".")[1]
	tags := sqliteTags(v)

	stmt := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (", tblName)
	for _, tag := range tags {
		tagElemes := strings.Split(tag.Tag, ",")
		switch len(tagElemes) {
		case 2:
			stmt = fmt.Sprintf("%s %s %s,", stmt, tagElemes[0], tagElemes[1])
		case 3:
			pkstmt := tagElemes[2]
			pkElems := strings.Split(pkstmt, "_")
			stmt = fmt.Sprintf("%s %s %s %s %s,", stmt, tagElemes[0], tagElemes[1], pkElems[0], pkElems[1])
		}
	}
	stmt = stmt[0 : len(stmt)-1] // remove last comma
	stmt = fmt.Sprintf(`%s )`, stmt)

	return stmt
}

func CreateInsertStmt[T DrawType](vals []T) string {
	stmt := "BEGIN TRANSACTIONS;"
	for _, val := range vals {
		s := createInsertStmt[T](&val)
		stmt = fmt.Sprintf("%s %s;", stmt, s)
	}
	stmt = fmt.Sprintf("%s COMMIT;", stmt)
	return stmt
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

// CreateTbl is an implementation of operations to create Table
func CreateTbl[T DrawType](db *sql.DB, creator CreateTblStmtFunc[T], typ *T) (sql.Result, error) {
	stmt := creator(typ)
	result, err := db.Exec(stmt)
	if err != nil {
		return result, fmt.Errorf("%w: %s", ErrDBCreateTbl, err.Error())
	}
	return result, nil
}
