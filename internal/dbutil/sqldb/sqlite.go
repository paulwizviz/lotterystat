// Package sqlite implementations operations related to the DevOps and
// the writing and reading of data to and from SQLite
package sqldb

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/paulwizviz/lotterystat/internal/csvdata"
)

var (

	// SQLiteCreateEuroDrawTbl is a default handler function to create a table in SQLite.
	//
	// Example statement
	// CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)
	SQLiteCreateEuroDrawTbl TblCreateHdlFunc[csvdata.EuroDraw] = func(db *sql.DB, tblName string, data csvdata.EuroDraw) error {
		stmtStr := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (", tblName)

		tags := data.SQLiteTags()
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

	// SQLiteWriteEuroDrawTbl is a default operation to insert object into sqlite table
	//
	// Example statement
	// INSERT INTO people (firstname, lastname) VALUES (?, ?)
	SQLiteWriteEuroDrawTbl TblWriterHdlFunc[csvdata.EuroDraw] = func(db *sql.DB, tblName string, data csvdata.EuroDraw) error {
		return nil
	}
)
