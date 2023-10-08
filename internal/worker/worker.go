package worker

import (
	"context"
	"database/sql"
	"paulwizviz/lotterystat/internal/euro"
	"paulwizviz/lotterystat/internal/sforl"
)

func InitalizeDB(ctx context.Context, db *sql.DB) error {
	err := euro.CreateTable(ctx, db)
	if err != nil {
		return err
	}
	err = sforl.CreateTable(ctx, db)
	if err != nil {
		return err
	}
	return nil
}

func PersistsDraw(ctx context.Context, db *sql.DB) error {
	err := euro.PersistsCSV(ctx, db, 3)
	if err != nil {
		return err
	}
	err = sforl.PersistsCSV(ctx, db, 3)
	if err != nil {
		return err
	}
	return nil
}
