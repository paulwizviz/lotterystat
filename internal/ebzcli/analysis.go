package ebzcli

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path"
	"paulwizviz/lotterystat/internal/dbutil"
	"paulwizviz/lotterystat/internal/sforl"

	"github.com/spf13/cobra"
)

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
			db, err := dbutil.SQLiteConnectFile(sqliteFile)
			if err != nil {
				log.Fatal(err)
			}
			bf, err := sforl.BallFreq(context.TODO(), db)
			if err != nil {
				log.Fatal(err)
			}
			sf, err := sforl.StarFreq(context.TODO(), db)
			if err != nil {
				log.Fatal(err)
			}

			pwd, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			outPath := path.Join(pwd, "tmp")
			starCSVFile := path.Join(outPath, "sforl-star.csv")
			ballCSVFile := path.Join(outPath, "sforl-ball.csv")
			f1, err := os.Create(starCSVFile)
			if err != nil {
				log.Fatal(err)
			}
			defer f1.Close()
			f2, err := os.Create(ballCSVFile)
			if err != nil {
				log.Fatal(err)
			}
			defer f2.Close()

			w1 := csv.NewWriter(f1)
			defer w1.Flush()

			starHeaders := []string{"Star", "Count"}
			var starData [][]string
			for _, r := range sf {
				d := []string{}
				d = append(d, fmt.Sprintf("%v", r.Star))
				d = append(d, fmt.Sprintf("%v", r.Count))
				starData = append(starData, d)
			}
			w1.Write(starHeaders)
			for _, row := range starData {
				w1.Write(row)
			}

			w2 := csv.NewWriter(f2)
			defer w2.Flush()
			ballHeaders := []string{"Ball", "Count"}
			var ballData [][]string
			for _, r := range bf {
				d := []string{}
				d = append(d, fmt.Sprintf("%v", r.Ball))
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

func sforlCmdInit() {
	sforlCmd.AddCommand(sforlFreqCmd)
}
