package main

import (
	"log"
	"os"
	"path"

	"github.com/spf13/cobra"
)

var dbPath string

var (
	initCmd = &cobra.Command{
		Use:   "init",
		Short: "init the analyzer",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	downloadCmd = &cobra.Command{
		Use:   "download",
		Short: "download from national lottery results",
		Run: func(cmd *cobra.Command, args []string) {
			home, err := os.UserHomeDir()
			if err != nil {
				log.Fatal(err)
			}

			dbPath := path.Join(home, ".lottery")
			_, err = os.Stat(dbPath)
			if os.IsNotExist(err) {
				err := os.Mkdir(dbPath, 0755)
				if err == nil {
					log.Println("db path created")
				}
			}
		},
	}

	rootCmd = &cobra.Command{
		Use:   "lotteryanalyzer",
		Short: "lotteryanalyzer download and analyze national lottery results",
	}
)

func init() {

	initCmd.PersistentFlags().StringVarP(&dbPath, "path", "p", "", "db store")
	rootCmd.AddCommand(initCmd)

	rootCmd.AddCommand(downloadCmd)
}

func execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	execute()
}
