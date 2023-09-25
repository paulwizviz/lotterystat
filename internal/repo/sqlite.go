package repo

import (
	"database/sql"
	"fmt"
	"paulwizviz/lotterystat/internal/euro"
	"reflect"
	"strings"
)

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

func sqliteCreateTblStmtStr[T drawType](v *T) string {
	t := reflect.TypeOf(v)
	tblName := strings.Split(fmt.Sprintf("%v", t), ".")[0][1:]
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

func sqliteCreateTbl(db *sql.DB, draw interface{}) error {
	switch d := draw.(type) {
	case *euro.Draw:
		_, err := db.Exec(sqliteCreateTblStmtStr(d))
		if err != nil {
			return fmt.Errorf("%w-%s", ErrCreateTbl, err.Error())
		}
		return nil
	case euro.Draw:
		_, err := db.Exec(sqliteCreateTblStmtStr(&d))
		if err != nil {
			return fmt.Errorf("%w-%s", ErrCreateTbl, err.Error())
		}
		return nil
	}
	return fmt.Errorf("%w-%s", ErrCreateTbl, "invalid draw types")
}

func sqliteInsertStmtStr[T drawType](v *T) string {
	t := reflect.TypeOf(v)
	tblName := strings.Split(fmt.Sprintf("%v", t), ".")[0][1:]
	tags := sqliteTags(v)

	stmt := fmt.Sprintf("INSERT INTO %s (", tblName)
	for _, tag := range tags {
		t1 := strings.Split(tag.Tag, ",")
		stmt = fmt.Sprintf("%s %s,", stmt, t1[0])
	}

	stmt = stmt[0 : len(stmt)-1]
	stmt = fmt.Sprintf("%s) VALUES (", stmt)
	for range tags {
		stmt = fmt.Sprintf("%s ?,", stmt)
	}
	stmt = stmt[0 : len(stmt)-1]
	stmt = fmt.Sprintf("%s )", stmt)
	return stmt
}

func sqlitePrepareWriteStmt[T drawType](db *sql.DB, v *T) (*sql.Stmt, error) {
	stmt, err := db.Prepare(sqliteInsertStmtStr(v))
	if err != nil {
		return nil, fmt.Errorf("%w-%s", ErrPrepareStmt, err.Error())
	}
	return stmt, nil
}

func sqlitePrepareTxWriteStmt[T drawType](tx *sql.Tx, v *T) (*sql.Stmt, error) {
	stmt, err := tx.Prepare(sqliteInsertStmtStr(v))
	if err != nil {
		return nil, fmt.Errorf("%w-%s", ErrPrepareStmt, err.Error())
	}
	return stmt, nil
}

func sqliteWriteTbl(db *sql.DB, draw interface{}) error {
	switch d := draw.(type) {
	case *euro.Draw:
		stmt, err := sqlitePrepareWriteStmt(db, d)
		if err != nil {
			return err
		}
		_, err = stmt.Exec(d.DrawDate, d.DayOfWeek, d.Ball1, d.Ball2, d.Ball3, d.Ball4, d.Ball5, d.LS1, d.LS2, d.UKMarker, d.EuroMarker, d.DrawNo)
		if err != nil {
			return fmt.Errorf("%w-%s", ErrWriteTbl, err.Error())
		}
		return nil
	case euro.Draw:
		stmt, err := sqlitePrepareWriteStmt(db, &d)
		if err != nil {
			return err
		}
		_, err = stmt.Exec(d.DrawDate, d.DayOfWeek, d.Ball1, d.Ball2, d.Ball3, d.Ball4, d.Ball5, d.LS1, d.LS2, d.UKMarker, d.EuroMarker, d.DrawNo)
		if err != nil {
			return fmt.Errorf("%w-%s", ErrWriteTbl, err.Error())
		}
		return nil
	}
	return fmt.Errorf("%w-%s", ErrWriteTbl, "invalid draw types")
}

func sqliteWriteTxTbl(tx *sql.Tx, draw interface{}) error {
	switch d := draw.(type) {
	case *euro.Draw:
		stmt, err := sqlitePrepareTxWriteStmt(tx, d)
		if err != nil {
			return err
		}
		_, err = stmt.Exec(d.DrawDate, d.DayOfWeek, d.Ball1, d.Ball2, d.Ball3, d.Ball4, d.Ball5, d.LS1, d.LS2, d.UKMarker, d.EuroMarker, d.DrawNo)
		if err != nil {
			return fmt.Errorf("%w-%s", ErrWriteTbl, err.Error())
		}
		return nil
	case euro.Draw:
		stmt, err := sqlitePrepareTxWriteStmt(tx, &d)
		if err != nil {
			return err
		}
		_, err = stmt.Exec(d.DrawDate, d.DayOfWeek, d.Ball1, d.Ball2, d.Ball3, d.Ball4, d.Ball5, d.LS1, d.LS2, d.UKMarker, d.EuroMarker, d.DrawNo)
		if err != nil {
			return fmt.Errorf("%w-%s", ErrWriteTbl, err.Error())
		}
		return nil
	}
	return fmt.Errorf("%w-%s", ErrWriteTbl, "invalid draw types")
}
