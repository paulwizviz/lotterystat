package sqlops_test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/paulwizviz/lotterystat/internal/sqlops"

	_ "github.com/mattn/go-sqlite3"
)

type tableColumn struct {
	OID       int
	Name      string
	DataType  string
	NotNull   int
	PK        int
	DFltValue sql.NullString
}

type table struct {
	Name    string
	Columns []tableColumn
}

// listTables is an operation to get names of tables
func listTables(db *sql.DB) ([]string, error) {
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

func listColumns(db *sql.DB, tableName string) (table, error) {

	tbl := table{
		Name: tableName,
	}

	query := fmt.Sprintf("PRAGMA table_info(%s)", tableName)
	rows, err := db.Query(query)
	if err != nil {
		return table{}, err
	}
	defer rows.Close()

	cols := []tableColumn{}
	for rows.Next() {
		col := tableColumn{}
		var oid int
		var name, dataType string
		var notNull, pk int
		var dfltValue sql.NullString

		err := rows.Scan(&oid, &name, &dataType, &notNull, &dfltValue, &pk)
		if err != nil {
			return table{}, err
		}
		col.OID = oid
		col.Name = name
		col.DataType = dataType
		col.NotNull = notNull
		col.PK = pk
		col.DFltValue = dfltValue
		cols = append(cols, col)
	}
	tbl.Columns = cols

	return tbl, nil
}

var (
	createTblSuccessCases = func(t *testing.T) {
		testcases := []struct {
			name          string
			input         sqlops.TblCreator
			wantTableName string
			wantTable     table
			wantErr       error
		}{
			{
				name: "Case 1",
				input: func(ctx context.Context, tx *sql.Tx) error {
					_, err := tx.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS draw(id INTEGER PRIMARY KEY, ball1 INTEGER)")
					if err != nil {
						return err
					}
					return nil
				},
				wantTableName: "draw",
				wantTable: table{
					Name: "draw",
					Columns: []tableColumn{
						{
							OID:      0,
							Name:     "id",
							DataType: "INTEGER",
							NotNull:  0,
							PK:       1,
							DFltValue: sql.NullString{
								Valid: false,
							},
						},
						{
							OID:      1,
							Name:     "ball1",
							DataType: "INTEGER",
							NotNull:  0,
							PK:       0,
							DFltValue: sql.NullString{
								Valid: false,
							},
						},
					},
				},
				wantErr: nil,
			},
		}

		for _, tc := range testcases {
			t.Run(tc.name, func(t *testing.T) {
				db, _ := sqlops.NewSQLiteMem()
				defer db.Close()
				gotErr := sqlops.CreateTables(context.TODO(), db, tc.input)
				if !errors.Is(gotErr, tc.wantErr) {
					t.Fatalf("Umatch error. What: %v Got: %v", tc.wantErr, gotErr)
				}
				tblnames, _ := listTables(db)
				if tblnames[0] != tc.wantTableName {
					t.Fatalf("Unmatch table name. Want: %v Got: %v", tc.wantTableName, tblnames[0])
				}
				gotTbl, _ := listColumns(db, tblnames[0])
				if !reflect.DeepEqual(gotTbl, tc.wantTable) {
					t.Fatalf("Unmatch tables. Want: %v Got: %v", tc.wantTable, gotTbl)
				}
			})
		}
	}
)

func TestCreateTable(t *testing.T) {
	t.Run("Success", createTblSuccessCases)
}

// CREATE TABLE IF NOT EXISTS draw(id INTEGER PRIMARY KEY, ball1 INTEGER

type draw struct {
	ID    int
	Ball1 int
}

func Example_writeTable() {
	db, _ := sqlops.NewSQLiteMem()
	defer db.Close()

	err := sqlops.CreateTables(context.TODO(), db, func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS draw(id INTEGER PRIMARY KEY, ball1 INTEGER)")
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println("create table", err)
	}

	data := draw{
		Ball1: 1,
	}
	err = sqlops.Writer(context.TODO(), db, `INSERT INTO draw (ball1) VALUES($1)`, []any{data}, func(ctx context.Context, stmt *sql.Stmt, data any) error {
		d, ok := data.(draw)
		if !ok {
			return fmt.Errorf("unable to cast to appropriate type")
		}
		_, err := stmt.ExecContext(ctx, d.Ball1)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query("SELECT * FROM draw")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var ball1 int
		rows.Scan(&id, &ball1)
		fmt.Println(id, ball1)
	}

	// Output:
	// 1 1
}

func Example_queryTable() {
	db, _ := sqlops.NewSQLiteMem()
	defer db.Close()

	err := sqlops.CreateTables(context.TODO(), db, func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS draw(id INTEGER PRIMARY KEY, ball1 INTEGER)")
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println("create table", err)
	}

	data := draw{
		Ball1: 1,
	}
	err = sqlops.Writer(context.TODO(), db, `INSERT INTO draw (ball1) VALUES($1)`, []any{data}, func(ctx context.Context, stmt *sql.Stmt, data any) error {
		d, ok := data.(draw)
		if !ok {
			return fmt.Errorf("unable to cast to appropriate type")
		}
		_, err := stmt.ExecContext(ctx, d.Ball1)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	result, err := sqlops.Query(context.TODO(), db, func(rows *sql.Rows) (any, error) {
		d := draw{}
		err := rows.Scan(&d.ID, &d.Ball1)
		if err != nil {
			return nil, err
		}
		return d, nil
	}, `SELECT * FROM draw`)

	fmt.Println(result)

	// Output:
	// [{1 1}]
}
