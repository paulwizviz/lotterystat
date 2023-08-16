// Package draw implements datasvc draws
package draw

import (
	"errors"
	"fmt"
	"reflect"
)

// Database related operations

var (
	ErrDBStatement = errors.New("statement error")
)

func CreateInsertStmt(val any) (string, error) {
	switch v := val.(type) {
	case []Euro:
		stmt := "BEGIN TRANSACTION;"
		for _, e := range v {
			stmt = fmt.Sprintf(`%s;\n%s;`, createInsertStmt(&e), stmt)
		}
		return stmt + "\nCOMMIT;", nil
	case []Set4Life:
		stmt := "BEGIN TRANSACTION;"
		for _, e := range v {
			stmt = fmt.Sprintf(`%s;\n%s;`, createInsertStmt(&e), stmt)
		}
		return stmt + "\nCOMMIT;", nil
	default:
		return "", fmt.Errorf("invalid type: %v", reflect.TypeOf(v))
	}
}

type DrawType interface {
	Euro | Set4Life
}
