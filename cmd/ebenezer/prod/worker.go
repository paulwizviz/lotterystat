package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"paulwizviz/lotterystat/internal/euro"
	"paulwizviz/lotterystat/internal/sforl"
	"sync"
)

func initalizeDB(ctx context.Context, db *sql.DB) error {
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

func persistsDraw(ctx context.Context, db *sql.DB) error {

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

func euroMatch(ctx context.Context, bet string, output string, db *sql.DB) error {

	if !euro.IsValidBet(bet) {
		return fmt.Errorf("can't bet")
	}
	b, err := euro.ProcessBetArg(bet)
	if err != nil {
		return fmt.Errorf("can't bet")
	}

	if output == "" {
		return fmt.Errorf("no file name")
	}
	f, err := os.Create(output)
	if err != nil {
		return fmt.Errorf("unable to create output file")
	}
	defer f.Close()
	w := csv.NewWriter(f)
	defer w.Flush()

	bets := []euro.Bet{b}
	mbs, err := euro.MatchBets(ctx, db, bets)
	if err != nil {
		return err
	}

	headers := []string{"Bet", "Draw", "Match Balls", "Lucky Stars"}
	var data [][]string
	for _, mb := range mbs {
		var d []string
		d = append(d, fmt.Sprintf("%v", mb.Bet))
		d = append(d, fmt.Sprintf("{%d,%d,%d,%d,%d,%d,%d}", mb.Draw.Ball1, mb.Draw.Ball2, mb.Draw.Ball3, mb.Draw.Ball4, mb.Draw.Ball5, mb.Draw.LS1, mb.Draw.LS2))
		d = append(d, fmt.Sprintf("%d", mb.Balls))
		d = append(d, fmt.Sprintf("%d", mb.LuckyStars))
		data = append(data, d)
	}
	w.Write(headers)
	for _, row := range data {
		w.Write(row)
	}
	return nil
}

func euroBallsFreq(ctx context.Context, output string, db *sql.DB) error {
	balls := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50}
	result, err := euro.CountBalls(ctx, db, balls)
	if err != nil {
		return fmt.Errorf("can't count balls")
	}
	if output == "" {
		return fmt.Errorf("no file name")
	}
	f, err := os.Create(output)
	if err != nil {
		return fmt.Errorf("unable to create output file")
	}
	defer f.Close()
	w := csv.NewWriter(f)
	defer w.Flush()
	headers := []string{"Ball", "Count"}
	var data [][]string
	for _, r := range result {
		d := []string{}
		d = append(d, fmt.Sprintf("%v", r.Ball))
		d = append(d, fmt.Sprintf("%v", r.Count))
		data = append(data, d)
	}
	w.Write(headers)
	for _, row := range data {
		w.Write(row)
	}
	return nil
}

func euroStarsFreq(ctx context.Context, output string, db *sql.DB) error {
	stars := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	result, err := euro.CountStars(ctx, db, stars)
	if err != nil {
		return fmt.Errorf("unable to count stars")
	}
	if output == "" {
		return fmt.Errorf("no file name")
	}
	f, err := os.Create(output)
	if err != nil {
		return fmt.Errorf("unable to create output file")
	}
	defer f.Close()
	w := csv.NewWriter(f)
	defer w.Flush()
	headers := []string{"Star", "Count"}
	var data [][]string
	for _, r := range result {
		d := []string{}
		d = append(d, fmt.Sprintf("%v", r.Star))
		d = append(d, fmt.Sprintf("%v", r.Count))
		data = append(data, d)
	}
	w.Write(headers)
	for _, row := range data {
		w.Write(row)
	}
	return nil
}

func sForLMatch(ctx context.Context, arg string, db *sql.DB) error {

	if !sforl.IsValidBet(arg) {
		return fmt.Errorf("can't bet")
	}
	b, err := sforl.ProcessBetArg(arg)
	if err != nil {
		return fmt.Errorf("can't bet")
	}
	bets := []sforl.Bet{b}
	mbs, err := sforl.MatchBets(ctx, db, bets)
	if err != nil {
		return err
	}

	for _, mb := range mbs {
		fmt.Printf("Bet: %v Draw: %v Match Balls: %v Life ball: %v\n", mb.Bet, fmt.Sprintf("{%d,%d,%d,%d,%d,%d}", mb.Draw.Ball1, mb.Draw.Ball2, mb.Draw.Ball3, mb.Draw.Ball4, mb.Draw.Ball5, mb.Draw.LifeBall), mb.Balls, mb.LifeBall)
	}

	return nil
}

func sForLBallsFreq(ctx context.Context, output string, db *sql.DB) error {
	balls := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47}
	result, err := sforl.CountBalls(ctx, db, balls)
	if err != nil {
		return fmt.Errorf("can't count balls")
	}
	if output == "" {
		return fmt.Errorf("no file name")
	}
	f, err := os.Create(output)
	if err != nil {
		return fmt.Errorf("unable to create output file")
	}
	defer f.Close()
	w := csv.NewWriter(f)
	defer w.Flush()
	headers := []string{"Ball", "Count"}
	var data [][]string
	for _, r := range result {
		d := []string{}
		d = append(d, fmt.Sprintf("%v", r.Ball))
		d = append(d, fmt.Sprintf("%v", r.Count))
		data = append(data, d)
	}
	w.Write(headers)
	for _, row := range data {
		w.Write(row)
	}
	return nil
}

func sForLLuckyBallFreq(ctx context.Context, output string, db *sql.DB) error {
	stars := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result, err := sforl.CountLuckyBall(ctx, db, stars)
	if err != nil {
		return fmt.Errorf("unable to count stars")
	}
	if output == "" {
		return fmt.Errorf("no file name")
	}
	f, err := os.Create(output)
	if err != nil {
		return fmt.Errorf("unable to create output file")
	}
	defer f.Close()
	w := csv.NewWriter(f)
	defer w.Flush()
	headers := []string{"Lucky ball", "Count"}
	var data [][]string
	for _, r := range result {
		d := []string{}
		d = append(d, fmt.Sprintf("%v", r.LuckyBall))
		d = append(d, fmt.Sprintf("%v", r.Count))
		data = append(data, d)
	}
	w.Write(headers)
	for _, row := range data {
		w.Write(row)
	}
	return nil
}
