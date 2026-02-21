package ebzcli

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/paulwizviz/lotterystat/internal/csvops"
	"github.com/paulwizviz/lotterystat/internal/ebzconfig"
	"github.com/paulwizviz/lotterystat/internal/sflife"
	"github.com/paulwizviz/lotterystat/internal/sqlops"
	"github.com/spf13/cobra"
)

var (
	sflifeFile string
)

func init() {
	sflifeCmd.AddCommand(sflifePersistsCmd)
	sflifePersistsCmd.Flags().StringVarP(&sflifeFile, "file", "f", "", "Set For Life CSV file to persist")
}

var sflifeCmd = &cobra.Command{
	Use:   "sflife",
	Short: "sflife is a subcommand related to Set For Life draws",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var sflifePersistsCmd = &cobra.Command{
	Use:   "persists",
	Short: "persist Set For Life csv file",
	Run: func(cmd *cobra.Command, args []string) {
		if sflifeFile == "" {
			cmd.Help()
			return
		}

		f, err := os.Open(sflifeFile)
		if err != nil {
			log.Fatalf("unable to open file %s: %v", sflifeFile, err)
		}
		defer f.Close()

		ctx := context.Background()
		recs := csvops.ExtractRec(ctx, f)
		drawChans := sflife.ProcessCSV(recs, 5)

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
			if err := sflife.PersistsDraw(ctx, db, dc.Draw); err != nil {
				log.Printf("unable to persist draw %v: %v", dc.Draw.DrawNo, err)
			}
		}
		fmt.Printf("Successfully processed %d records\n", len(drawChans))
	},
}
