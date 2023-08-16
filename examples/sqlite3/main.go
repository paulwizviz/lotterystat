package main

import (
	"database/sql"
	"log"
	"log/slog"
	"os"
	"path"
	"paulwizviz/lotterystat/internal/draw"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	hdl := slog.NewJSONHandler(os.Stdout, opts)
	slog.SetDefault(slog.New(hdl))

	homePath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dbPath := path.Join(homePath, "tmp", "test.db")

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var createStmtFn draw.CreateTblStmtfuncL[draw.Euro] = draw.CreateTblStmt
	var createTblFn draw.CreateTblFuncL[draw.Euro] = draw.CreateTbl

	createTblFn.Create(db, createStmtFn, &draw.Euro{})

}
