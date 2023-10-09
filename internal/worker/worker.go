package worker

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"paulwizviz/lotterystat/internal/euro"
	"paulwizviz/lotterystat/internal/sforl"
	"sync"
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

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		err := euro.PersistsCSV(ctx, db, 3)
		if err != nil {
			log.Println(err)
			return
		}
	}()

	go func() {
		defer wg.Done()
		err := sforl.PersistsCSV(ctx, db, 3)
		if err != nil {
			log.Println(err)
			return
		}
	}()

	wg.Wait()
	return nil
}

func ProcessEuroBetArg(ctx context.Context, arg string, db *sql.DB) error {

	if !euro.IsValidBet(arg) {
		return fmt.Errorf("can't bet")
	}
	b, err := euro.ProcessBetArg(arg)
	if err != nil {
		return fmt.Errorf("can't bet")
	}
	bets := []euro.Bet{b}
	mbs, err := euro.MatchBets(ctx, db, bets)
	if err != nil {
		return err
	}

	for _, mb := range mbs {
		fmt.Printf("Bet: %v Draw: %v Match Balls: %v Lucky Stars: %v\n", mb.Bet, fmt.Sprintf("{%d,%d,%d,%d,%d,%d,%d}", mb.Draw.Ball1, mb.Draw.Ball2, mb.Draw.Ball3, mb.Draw.Ball4, mb.Draw.Ball5, mb.Draw.LS1, mb.Draw.LS2), mb.Balls, mb.LuckyStars)
	}

	return nil
}
