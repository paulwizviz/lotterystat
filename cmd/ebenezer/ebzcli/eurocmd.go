package main

import (
	"log"

	"github.com/paulwizviz/lotterystat/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	euroCmd = &cobra.Command{
		Use:   "euro",
		Short: "A subcommand to process Euro draws",
		Run: func(cmd *cobra.Command, args []string) {
			p := viper.GetString(config.SQLitePathKey)
			f := viper.GetString(config.SQLiteFileKey)
			log.Printf("Path: %s File: %s", p, f)
		},
	}
)
