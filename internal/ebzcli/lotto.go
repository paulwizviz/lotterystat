package ebzcli

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/paulwizviz/lotterystat/internal/csvops"
	"github.com/paulwizviz/lotterystat/internal/ebzconfig"
	"github.com/paulwizviz/lotterystat/internal/lotto"
	"github.com/paulwizviz/lotterystat/internal/sqlops"
	"github.com/spf13/cobra"
)

var (
	lottoFile string
)

func init() {
	lottoCmd.AddCommand(lottoPersistsCmd)
	lottoPersistsCmd.Flags().StringVarP(&lottoFile, "file", "f", "", "Lotto CSV file to persist")
}

var lottoCmd = &cobra.Command{
	Use:   "lotto",
	Short: "lotto is a subcommand related to Lotto draws",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var lottoPersistsCmd = &cobra.Command{
	Use:   "persists",
	Short: "persist Lotto csv file",
	Run: func(cmd *cobra.Command, args []string) {
		if lottoFile == "" {
			cmd.Help()
			return
		}

		f, err := os.Open(lottoFile)
		if err != nil {
			log.Fatalf("unable to open file %s: %v", lottoFile, err)
		}
		defer f.Close()

		ctx := context.Background()
		recs := csvops.ExtractRec(ctx, f)
		drawChans := lotto.ProcessCSV(recs, 5)

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
			if err := lotto.PersistsDraw(ctx, db, dc.Draw); err != nil {
				log.Printf("unable to persist draw %v: %v", dc.Draw.DrawNo, err)
			}
		}
		fmt.Printf("Successfully processed %d records\n", len(drawChans))
	},
}
