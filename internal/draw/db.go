package draw

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
)

var (
	ErrUnableToCreateTable = errors.New("unable to create table")
	ErrDuplicateEntry      = errors.New("dublicate entry error")
)

type drawType interface {
	Euro | Set4Life
}

type structTag struct {
	FieldName string
	Tag       string
}

func sqliteTags[T drawType](typ *T) []structTag {
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

func createTblStmt[T drawType](v *T) string {

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

func insertTblStmt[T drawType](val *T) string {
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
