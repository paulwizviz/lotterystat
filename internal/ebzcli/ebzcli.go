package ebzcli

import (
	"log"
	"os"
	"path"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ebz",
	Short: "ebz is a cli app to help you analyze UK National Lottery.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	dbCmdSetup()
	sforlCmdSetup()
	euroCmdSetup()
	rootCmd.AddCommand(dbCmd)
	rootCmd.AddCommand(sforlCmd)
	rootCmd.AddCommand(euroCmd)
}

func Execute() error {
	sqliteFile = func() string {
		sFile := os.Getenv("SQLITE_DB")
		if sFile == "" {
			pwd, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			sFile = path.Join(pwd, "dbfiles", "sqlite", "data.db")
		}
		return sFile
	}()
	return rootCmd.Execute()
}
