package ebzcli

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/paulwizviz/lotterystat/internal/dbutil"
	"github.com/paulwizviz/lotterystat/internal/euro"
	"github.com/paulwizviz/lotterystat/internal/sforl"
	"github.com/paulwizviz/lotterystat/internal/tball"

	"github.com/spf13/cobra"
)

// Set for life
var (
	sforlCmd = &cobra.Command{
		Use:   "sforl",
		Short: "sforl lottery data",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	sforlFreqCmd = &cobra.Command{
		Use:   "freq",
		Short: "Frequency analsysis",
		Run: func(cmd *cobra.Command, args []string) {
			// Obtaining DB connection
			db, err := dbutil.SQLiteConnectFile(sqliteFile)
			if err != nil {
				log.Fatal(err)
			}

			// Path to location for CSV files
			pwd, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			outPath := path.Join(pwd, "tmp")

			// Star frequencies.
			sf, err := sforl.StarFreq(context.TODO(), db)
			if err != nil {
				log.Fatal(err)
			}

			starCSVFile := path.Join(outPath, "sforl-star.csv")
			f1, err := os.Create(starCSVFile)
			if err != nil {
				log.Fatal(err)
			}
			defer f1.Close()
			w1 := csv.NewWriter(f1)
			defer w1.Flush()

			starHeaders := []string{"Star", "Count"}
			var starData [][]string
			for _, r := range sf {
				d := []string{}
				d = append(d, fmt.Sprintf("%v", r.Num))
				d = append(d, fmt.Sprintf("%v", r.Count))
				starData = append(starData, d)
			}
			w1.Write(starHeaders)
			for _, row := range starData {
				w1.Write(row)
			}

			// Main balls frequencies
			bf, err := sforl.MainFreq(context.TODO(), db)
			if err != nil {
				log.Fatal(err)
			}

			ballCSVFile := path.Join(outPath, "sforl-ball.csv")
			f2, err := os.Create(ballCSVFile)
			if err != nil {
				log.Fatal(err)
			}
			defer f2.Close()
			w2 := csv.NewWriter(f2)
			defer w2.Flush()
			ballHeaders := []string{"Ball", "Count"}
			var ballData [][]string
			for _, r := range bf {
				d := []string{}
				d = append(d, fmt.Sprintf("%v", r.Num))
				d = append(d, fmt.Sprintf("%v", r.Count))
				ballData = append(ballData, d)
			}
			w2.Write(ballHeaders)
			for _, row := range ballData {
				w2.Write(row)
			}

			// Two combination frequencies
			numworkers := 4
			comboFreq := sforl.TwoMainComboFreq(context.TODO(), db, numworkers)

			comboCSVFile := path.Join(outPath, "sforl-combo.csv")
			f3, err := os.Create(comboCSVFile)
			if err != nil {
				log.Fatal(err)
			}
			defer f3.Close()
			w3 := csv.NewWriter(f3)
			defer w3.Flush()

			comboHeaders := []string{"Combination", "Count"}
			var comboData [][]string
			for _, cf := range comboFreq {
				d := []string{}
				if cf.Count != 0 {
					d = append(d, fmt.Sprintf("%v", cf.Combo))
					d = append(d, fmt.Sprintf("%v", cf.Count))
					comboData = append(comboData, d)
				}
			}
			w3.Write(comboHeaders)
			for _, row := range comboData {
				w3.Write(row)
			}
		},
	}
)

func sforlCmdSetup() {
	sforlCmd.AddCommand(sforlFreqCmd)
}

// Euro
var (
	euroCmd = &cobra.Command{
		Use:   "euro",
		Short: "euro lottery data",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	euroFreqCmd = &cobra.Command{
		Use:   "freq",
		Short: "Frequency analsysis",
		Run: func(cmd *cobra.Command, args []string) {
			// Obtaining DB connection
			db, err := dbutil.SQLiteConnectFile(sqliteFile)
			if err != nil {
				log.Fatal(err)
			}

			// Path to location for CSV files
			pwd, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			outPath := path.Join(pwd, "tmp")

			// Star frequencies.
			lf, err := euro.LuckyFreq(context.TODO(), db)
			if err != nil {
				log.Fatal(err)
			}

			luckyCSVFile := path.Join(outPath, "euro-lucky.csv")
			f1, err := os.Create(luckyCSVFile)
			if err != nil {
				log.Fatal(err)
			}
			defer f1.Close()
			w1 := csv.NewWriter(f1)
			defer w1.Flush()

			luckyHeaders := []string{"Lucky Ball", "Count"}
			var starData [][]string
			for _, r := range lf {
				d := []string{}
				d = append(d, fmt.Sprintf("%v", r.Num))
				d = append(d, fmt.Sprintf("%v", r.Count))
				starData = append(starData, d)
			}
			w1.Write(luckyHeaders)
			for _, row := range starData {
				w1.Write(row)
			}

			// Main balls frequencies
			bf, err := euro.MainFreq(context.TODO(), db)
			if err != nil {
				log.Fatal(err)
			}

			ballCSVFile := path.Join(outPath, "euro-ball.csv")
			f2, err := os.Create(ballCSVFile)
			if err != nil {
				log.Fatal(err)
			}
			defer f2.Close()
			w2 := csv.NewWriter(f2)
			defer w2.Flush()
			ballHeaders := []string{"Ball", "Count"}
			var ballData [][]string
			for _, r := range bf {
				d := []string{}
				d = append(d, fmt.Sprintf("%v", r.Num))
				d = append(d, fmt.Sprintf("%v", r.Count))
				ballData = append(ballData, d)
			}
			w2.Write(ballHeaders)
			for _, row := range ballData {
				w2.Write(row)
			}
		},
	}
)

func euroCmdSetup() {
	euroCmd.AddCommand(euroFreqCmd)
}

// Tball
var (
	tballCmd = &cobra.Command{
		Use:   "tball",
		Short: "thunder ball lottery data",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	tballFreqCmd = &cobra.Command{
		Use:   "freq",
		Short: "Frequency analsysis",
		Run: func(cmd *cobra.Command, args []string) {
			// Obtaining DB connection
			db, err := dbutil.SQLiteConnectFile(sqliteFile)
			if err != nil {
				log.Fatal(err)
			}

			// Path to location for CSV files
			pwd, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			outPath := path.Join(pwd, "tmp")

			// tball frequencies.
			lf, err := tball.TballFreq(context.TODO(), db)
			if err != nil {
				log.Fatal(err)
			}

			tballCSVFile := path.Join(outPath, "thunderball-tball.csv")
			f1, err := os.Create(tballCSVFile)
			if err != nil {
				log.Fatal(err)
			}
			defer f1.Close()
			w1 := csv.NewWriter(f1)
			defer w1.Flush()

			tballHeaders := []string{"Thuderball Ball", "Count"}
			var tballData [][]string
			for _, r := range lf {
				d := []string{}
				d = append(d, fmt.Sprintf("%v", r.Num))
				d = append(d, fmt.Sprintf("%v", r.Count))
				tballData = append(tballData, d)
			}
			w1.Write(tballHeaders)
			for _, row := range tballData {
				w1.Write(row)
			}

			// Main balls frequencies
			bf, err := tball.MainFreq(context.TODO(), db)
			if err != nil {
				log.Fatal(err)
			}

			ballCSVFile := path.Join(outPath, "thunderball-ball.csv")
			f2, err := os.Create(ballCSVFile)
			if err != nil {
				log.Fatal(err)
			}
			defer f2.Close()
			w2 := csv.NewWriter(f2)
			defer w2.Flush()
			ballHeaders := []string{"Ball", "Count"}
			var ballData [][]string
			for _, r := range bf {
				d := []string{}
				d = append(d, fmt.Sprintf("%v", r.Num))
				d = append(d, fmt.Sprintf("%v", r.Count))
				ballData = append(ballData, d)
			}
			w2.Write(ballHeaders)
			for _, row := range ballData {
				w2.Write(row)
			}
		},
	}
)

func tballCmdSetup() {
	tballCmd.AddCommand(tballFreqCmd)
}
