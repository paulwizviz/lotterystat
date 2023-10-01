package main

import (
	"log"
	"paulwizviz/lotterystat/internal/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "ebz",
	Short: "ebz is a cli base app to analyze UK national lottery stat",
	Run: func(cmd *cobra.Command, args []string) {
		dbConn := viper.Get(config.DBConnKey)
		log.Println(dbConn)
	},
}

func init() {

}

func Execute() error {
	return rootCmd.Execute()
}
