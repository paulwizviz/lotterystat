package ebzcli

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/paulwizviz/lotterystat/internal/csvops"
	"github.com/paulwizviz/lotterystat/internal/ebzconfig"
	"github.com/paulwizviz/lotterystat/internal/euro"
	"github.com/paulwizviz/lotterystat/internal/sqlops"
	"github.com/spf13/cobra"
)

var (
	euroFile string
)

func init() {
	euroCmd.AddCommand(euroPersistsCmd)
	euroPersistsCmd.Flags().StringVarP(&euroFile, "file", "f", "", "EuroMillions CSV file to persist")
}

var euroCmd = &cobra.Command{
	Use:   "euro",
	Short: "euro is a subcommand related to EuroMillions draws",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var euroPersistsCmd = &cobra.Command{
	Use:   "persists",
	Short: "persist EuroMillions csv file",
	Run: func(cmd *cobra.Command, args []string) {
		if euroFile == "" {
			cmd.Help()
			return
		}

		f, err := os.Open(euroFile)
		if err != nil {
			log.Fatalf("unable to open file %s: %v", euroFile, err)
		}
		defer f.Close()

		ctx := context.Background()
		recs := csvops.ExtractRec(ctx, f)
		drawChans := euro.ProcessCSV(recs, 5)

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
			if err := euro.PersistsDraw(ctx, db, dc.Draw); err != nil {
				log.Printf("unable to persist draw %v: %v", dc.Draw.DrawNo, err)
			}
		}
		fmt.Printf("Successfully processed %d records\n", len(drawChans))
	},
}
