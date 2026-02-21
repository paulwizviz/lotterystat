package ebzcli

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/paulwizviz/lotterystat/internal/csvops"
	"github.com/paulwizviz/lotterystat/internal/ebzconfig"
	"github.com/paulwizviz/lotterystat/internal/sqlops"
	"github.com/paulwizviz/lotterystat/internal/tball"
	"github.com/spf13/cobra"
)

var (
	tballFile string
)

func init() {
	tballCmd.AddCommand(tballPersistsCmd)
	tballPersistsCmd.Flags().StringVarP(&tballFile, "file", "f", "", "Thunderball CSV file to persist")
}

var tballCmd = &cobra.Command{
	Use:   "tball",
	Short: "tball is a subcommand",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var tballPersistsCmd = &cobra.Command{
	Use:   "persists",
	Short: "persist thunderball csv file",
	Run: func(cmd *cobra.Command, args []string) {
		if tballFile == "" {
			cmd.Help()
			return
		}

		f, err := os.Open(tballFile)
		if err != nil {
			log.Fatalf("unable to open file %s: %v", tballFile, err)
		}
		defer f.Close()

		ctx := context.Background()
		recs := csvops.ExtractRec(ctx, f)
		drawChans := tball.ProcessCSV(recs, 5)

		db, err := sqlops.NewSQLiteFile(ebzconfig.AppConfig.DatabasePath)
		if err != nil {
			log.Fatalf("unable to open database: %v", err)
		}
		defer db.Close()

		for _, dc := range drawChans {
			if dc.Err != nil {
				log.Printf("skipping record due to error: %v", dc.Err)
				continue
			}
			if err := tball.PersistsDraw(ctx, db, dc.Draw); err != nil {
				log.Printf("unable to persist draw %v: %v", dc.Draw.DrawNo, err)
			}
		}
		fmt.Printf("Successfully processed %d records\n", len(drawChans))
	},
}
