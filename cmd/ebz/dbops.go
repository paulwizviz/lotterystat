package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"paulwizviz/lotterystat/internal/dbutil"
	"paulwizviz/lotterystat/internal/euro"
	"paulwizviz/lotterystat/internal/sforl"

	"github.com/spf13/cobra"

	_ "github.com/mattn/go-sqlite3"
)

var (
	sqliteFile string

	dbCmd = &cobra.Command{
		Use:   "db",
		Short: "db calls",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	dbinitCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialise DB",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	sqliteCmd = &cobra.Command{
		Use:   "sqlite",
		Short: "sqlite db",
		Run: func(cmd *cobra.Command, args []string) {
			if sqliteFile == "" {
				fmt.Println("Initialising SQLite db")
				pwd, err := os.Getwd()
				if err != nil {
					log.Fatal(err)
				}
				sqliteFile = path.Join(pwd, "dbfiles", "sqlite", "data.db")
			}
			err := initSQLiteDB(sqliteFile)
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	psqlCmd = &cobra.Command{
		Use:   "psql",
		Short: "psql db",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Initialise Postgres db")
		},
	}
)

func initDBCmd() {
	dbCmd.AddCommand(dbinitCmd)
	dbinitCmd.AddCommand(sqliteCmd)
	sqliteCmd.Flags().StringVarP(&sqliteFile, "file", "f", "", "SQLite file")
	dbinitCmd.AddCommand(psqlCmd)
}

func initSQLiteDB(file string) error {
	db, err := dbutil.SQLiteConnectFile(file)
	if err != nil {
		return err
	}
	err = sforl.CreateSQLiteTable(context.TODO(), db)
	if err != nil {
		return err
	}
	err = euro.CreateSQLiteTable(context.TODO(), db)
	if err != nil {
		return err
	}
	return nil
}
