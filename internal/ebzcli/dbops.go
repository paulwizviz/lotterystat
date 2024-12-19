package ebzcli

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/paulwizviz/lotterystat/internal/dbutil"
	"github.com/paulwizviz/lotterystat/internal/euro"
	"github.com/paulwizviz/lotterystat/internal/sforl"
	"github.com/paulwizviz/lotterystat/internal/tball"

	"github.com/spf13/cobra"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var (
	dbCmd = &cobra.Command{
		Use:   "db",
		Short: "db calls",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	dbInitCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialise DB",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	dbPersistsCmd = &cobra.Command{
		Use:   "persists",
		Short: "Persist DB",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

var (
	sqliteFile string

	sqliteInitCmd = &cobra.Command{
		Use:   "sqlite",
		Short: "sqlite db",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("-- Initialising SQLite db --")
			err := initSQLiteDB(sqliteFile)
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	sqlitePersistsCmd = &cobra.Command{
		Use:   "sqlite",
		Short: "sqlite db",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("-- Persist to SQLite --")
			db, err := dbutil.SQLiteConnectFile(sqliteFile)
			if err != nil {
				log.Fatal(err)
			}
			err = euro.PersistsCSV(context.TODO(), db, 3)
			if err != nil {
				log.Fatal(err)
			}
			err = sforl.PersistsCSV(context.TODO(), db, 3)
			if err != nil {
				log.Fatal(err)
			}
			err = tball.PersistsCSV(context.TODO(), db, 3)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
)

var (
	username = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	host     = os.Getenv("POSTGRES_HOST")
	port     = os.Getenv("POSTGRES_PORT")
	dbName   = os.Getenv("POSTGRES_DBNAME")

	psqlInitCmd = &cobra.Command{
		Use:   "psql",
		Short: "psql db",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("-- Initialise Postgres --")
			p, err := strconv.Atoi(port)
			if err != nil {
				log.Fatal(err)
			}
			err = initPSQLDB(username, password, host, p, dbName)
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	psqlPersistsCmd = &cobra.Command{
		Use:   "psql",
		Short: "persists csv to psql db",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("-- Persist to PSQL --")
			p, err := strconv.Atoi(port)
			if err != nil {
				log.Fatal(err)
			}
			psqlDB, err := dbutil.PSQLConnect(username, password, host, p, dbName)
			if err != nil {
				log.Fatal(err)
			}
			sqliteDBFile := os.Getenv("SQLITE_DB")
			sqliteDB, err := dbutil.SQLiteConnectFile(sqliteDBFile)
			if err != nil {
				log.Fatal(err)
			}
			err = euro.DuplicateData(context.TODO(), sqliteDB, psqlDB, 4)
			if err != nil {
				log.Fatal(err)
			}
			err = sforl.DuplicateData(context.TODO(), sqliteDB, psqlDB, 4)
			if err != nil {
				log.Fatal(err)
			}
			err = tball.DuplicateData(context.TODO(), sqliteDB, psqlDB, 4)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
)

func dbCmdSetup() {

	// Configure Init commands
	dbInitCmd.AddCommand(sqliteInitCmd)
	sqliteInitCmd.Flags().StringVarP(&sqliteFile, "file", "f", "", "SQLite file")
	dbInitCmd.AddCommand(psqlInitCmd)

	// Configure Persists command
	dbPersistsCmd.AddCommand(sqlitePersistsCmd)
	sqlitePersistsCmd.Flags().StringVarP(&sqliteFile, "file", "f", "", "SQLite file")
	dbPersistsCmd.AddCommand(psqlPersistsCmd)

	// Add sub commands to db command root
	dbCmd.AddCommand(dbInitCmd)
	dbCmd.AddCommand(dbPersistsCmd)
}

func initSQLiteDB(file string) error {
	db, err := dbutil.SQLiteConnectFile(file)
	if err != nil {
		return err
	}
	defer db.Close()

	err = sforl.CreateSQLiteTable(context.TODO(), db)
	if err != nil {
		return err
	}
	err = euro.CreateSQLiteTable(context.TODO(), db)
	if err != nil {
		return err
	}
	err = tball.CreateSQLiteTable(context.TODO(), db)
	if err != nil {
		return err
	}
	return nil
}

func initPSQLDB(username string, password string, host string, port int, dbname string) error {
	connStmt := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)
	db, err := sql.Open("postgres", connStmt)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = euro.CreatePSQLTable(context.TODO(), db)
	if err != nil {
		return err
	}

	err = sforl.CreatePSQLTable(context.TODO(), db)
	if err != nil {
		return err
	}
	err = tball.CreatePSQLTable(context.TODO(), db)
	if err != nil {
		return err
	}
	return nil
}
