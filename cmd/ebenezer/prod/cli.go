package main

import (
	"context"
	"log"
	"paulwizviz/lotterystat/internal/worker"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ebz",
	Short: "ebz is a cli app to help you analyze UK National Lottery.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var matchCmd = &cobra.Command{
	Use:   "match",
	Short: "an ebz subcommand enabling users to match bet against draws",
	Long: `match is an ebz subcommand to enable users to match bets and draws
in a sequential ball by ball, extra numbers basis. For example, in the case of Euro million
lottery, a match against a user bet is as follows:

euro bet:  1,2, 3, 4, 5, 1, 12
euro draw: 1,3,10,20,30, 2, 12

The match is [1] and [12] for balls`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var freqCmd = &cobra.Command{
	Use:   "freq",
	Short: "an ebz subcommand enabling users to perform frequency analysis",
	Long: `freq is an ebz subcommand to enable users to determine frequency of
of numeric occurrence.

TO DO`,
}

var (
	euroBet      string
	euroMatchCmd = &cobra.Command{
		Use:   "euro",
		Args:  cobra.MaximumNArgs(0),
		Short: "comand to perform analysis of euro draws",
		Long:  `euro is an operation to match euro draws against bets presented by users`,
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

var (
	sforlBet      string
	sforlMatchCmd = &cobra.Command{
		Use:   "sforl",
		Args:  cobra.MaximumNArgs(0),
		Short: "comand to perform analysis of set for life draws",
		Long:  `sforl is an operation to match set for life draws against bets presented by users`,
		Run: func(cmd *cobra.Command, args []string) {
			if sforlBet == "" {
				cmd.Help()
				return
			}
			err := worker.ProcessSForLBetArg(context.TODO(), sforlBet, DB)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
)

func initMatchCmd() {
	matchCmd.AddCommand(euroMatchCmd)
	matchCmd.AddCommand(sforlMatchCmd)
}

func initEuroCmd() {
	euroMatchCmd.PersistentFlags().StringVarP(&euroBet, "bet", "b", "", `enter 5 numbers, 2 lucky stars in this format "[1-50],[1-50],[1-50],[1-50],[1-50],[1-12],[1-12]"`)
}

func initSForLCmd() {
	sforlMatchCmd.PersistentFlags().StringVarP(&sforlBet, "bet", "b", "", `enter 5 numbers, 2 lifeball in this format "[1-47],[1-47],[1-47],[1-47],[1-47],[1-10]"`)
}

func init() {
	initEuroCmd()
	initSForLCmd()
	initMatchCmd()
	rootCmd.AddCommand(freqCmd)
	rootCmd.AddCommand(matchCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
