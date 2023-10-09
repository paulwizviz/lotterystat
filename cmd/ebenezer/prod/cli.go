package main

import (
	"context"
	"log"
	"paulwizviz/lotterystat/internal/worker"

	"github.com/spf13/cobra"
)

var (
	euroBet string
)

var (
	rootCmd = &cobra.Command{
		Use:   "ebz",
		Short: "ebz is a cli app to help you analyze UK National Lottery statistics",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	euroCmd = &cobra.Command{
		Use:   "euro",
		Args:  cobra.MaximumNArgs(0),
		Short: "euro is a sub command of ebz app to perform analysis",
		Run: func(cmd *cobra.Command, args []string) {
			if euroBet == "" {
				cmd.Help()
				return
			}

			err := worker.ProcessEuroBetArg(context.TODO(), euroBet, DB)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
)

func initEuroCmd() {
	euroCmd.PersistentFlags().StringVarP(&euroBet, "bet", "b", "", `enter 5 numbers, 2 lucky stars in this format "[1-50],[1-50],[1-50],[1-50],[1-50],[1-12],[1-12]"`)
}

func init() {
	initEuroCmd()
	rootCmd.AddCommand(euroCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
