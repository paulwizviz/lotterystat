package sforl

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func dbConn() (*sql.DB, error) {
	return sql.Open("sqlite3", ":memory:")
}

func listSQLiteTables(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT name FROM sqlite_schema WHERE type='table' ORDER BY name")
	if err != nil {
		return nil, err
	}
	var result []string
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			fmt.Printf("Table name not found: %v", err)
		}
		result = append(result, name)
	}
	return result, nil
}

func Example_listSQLiteTable() {
	db, err := dbConn()
	if err != nil {
		fmt.Printf("DB not found: %v", err)
		return
	}
	defer db.Close()
	err = createSQLiteTable(context.TODO(), db)
	if err != nil {
		fmt.Printf("Create table error: %v", err)
		return
	}
	tbls, err := listSQLiteTables(db)
	if err != nil {
		fmt.Printf("Unable to list tables: %v", err)
		return
	}

	for _, tbl := range tbls {
		fmt.Println(tbl)
	}

	// Output:
	// set_for_life
}
