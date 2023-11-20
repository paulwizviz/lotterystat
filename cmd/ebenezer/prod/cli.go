package main

import (
	"context"
	"log"

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
	Short: "An ebz subcommand enabling users to match bet against draws",
	Long: `match is an ebz subcommand to enable users to match bets and draws
in a sequential ball by ball, extra numbers basis. For example, in the case of Euro million
lottery, a match against a user bet is as follows:

euro bet:  1,2, 3, 4, 5, 1, 12
euro draw: 1,3,10,20,30, 2, 12

The match is [1] and [12] for balls`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			cmd.Help()
		}
		cmd.Help()
	},
}

var freqCmd = &cobra.Command{
	Use:   "freq",
	Short: "A subcommand enabling users to perform frequency analysis",
	Long: `freq is an ebz subcommand to enable users to determine frequency of
of numeric occurrence.`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
		cmd.Help()
	},
}

var output string

var (
	euroBetFlag       string
	euroFreqStarsFlag bool
	euroFreqBallsFlag bool
	euroMatchCmd      = &cobra.Command{
		Use:   "euro",
		Args:  cobra.MaximumNArgs(0),
		Short: "A subcomand to perform analysis of euro draws",
		Long:  `euro is an operation to match euro draws against bets presented by users`,
		Run: func(cmd *cobra.Command, args []string) {
			if euroBetFlag == "" {
				cmd.Help()
				return
			}
			err := euroMatch(context.TODO(), euroBetFlag, output, DB)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
	euroFreqCmd = &cobra.Command{
		Use:   "euro",
		Args:  cobra.MaximumNArgs(0),
		Short: "A subcomand to perform analysis of euro draws",
		Long:  `euro is an operation to match euro draws against bets presented by users`,
		Run: func(cmd *cobra.Command, args []string) {
			if euroFreqStarsFlag {
				euroStarsFreq(context.TODO(), output, DB)
				return
			}
			if euroFreqBallsFlag {
				euroBallsFreq(context.TODO(), output, DB)
				return
			}
			cmd.Help()
		},
	}
)

var (
	sforlBetFlag       string
	sforlFreqLBFlag    bool
	sforlFreqBallsFlag bool
	sforlMatchCmd      = &cobra.Command{
		Use:   "sforl",
		Args:  cobra.MaximumNArgs(0),
		Short: "A subcomand to perform analysis of set for life draws",
		Long:  `sforl is an operation to match set for life draws against bets presented by users`,
		Run: func(cmd *cobra.Command, args []string) {
			if sforlBetFlag == "" {
				cmd.Help()
				return
			}
			err := sForLMatch(context.TODO(), sforlBetFlag, output, DB)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
	sforlFreqCmd = &cobra.Command{
		Use:   "sforl",
		Args:  cobra.MaximumNArgs(0),
		Short: "A subcomand to perform analysis of set for life draws",
		Long:  `sforl is an operation to match set for life draws against bets presented by users`,
		Run: func(cmd *cobra.Command, args []string) {
			if sforlFreqBallsFlag {
				err := sForLBallsFreq(context.TODO(), output, DB)
				if err != nil {
					log.Fatal(err)
				}
				return
			}
			if sforlFreqLBFlag {
				err := sForLLuckyBallFreq(context.TODO(), output, DB)
				if err != nil {
					log.Fatal(err)
				}
				return
			}
			cmd.Help()
		},
	}
)

func initEuroCmd() {
	euroMatchCmd.PersistentFlags().StringVarP(&euroBetFlag, "bet", "b", "", `enter 5 numbers, 2 lucky stars in this format "[1-50],[1-50],[1-50],[1-50],[1-50],[1-12],[1-12]"`)
	euroMatchCmd.PersistentFlags().StringVarP(&output, "output", "o", "", `where you like the result`)
	euroMatchCmd.MarkFlagRequired("output")
	euroFreqCmd.PersistentFlags().BoolVarP(&euroFreqBallsFlag, "balls", "b", false, `get frequency list for balls`)
	euroFreqCmd.PersistentFlags().BoolVarP(&euroFreqStarsFlag, "stars", "s", false, `get frequency list for stars`)
	euroFreqCmd.PersistentFlags().StringVarP(&output, "output", "o", "", `where we should output the result`)
	euroFreqCmd.MarkFlagRequired("output")
}

func initSForLCmd() {
	sforlMatchCmd.PersistentFlags().StringVarP(&sforlBetFlag, "bet", "b", "", `enter 5 numbers, 2 lifeball in this format "[1-47],[1-47],[1-47],[1-47],[1-47],[1-10]"`)
	sforlMatchCmd.PersistentFlags().StringVarP(&output, "output", "o", "", `where you like the result`)
	sforlMatchCmd.MarkFlagRequired("output")
	sforlFreqCmd.PersistentFlags().BoolVarP(&sforlFreqBallsFlag, "balls", "b", false, `get frequency list for balls`)
	sforlFreqCmd.PersistentFlags().BoolVarP(&sforlFreqLBFlag, "lucky", "l", false, `get frequency list for stars`)
	sforlFreqCmd.PersistentFlags().StringVarP(&output, "output", "o", "", `where we should output the result`)
	sforlFreqCmd.MarkFlagRequired("output")
}

func initMatchCmd() {
	matchCmd.AddCommand(euroMatchCmd)
	matchCmd.AddCommand(sforlMatchCmd)
}

func initFreqCmd() {
	freqCmd.AddCommand(euroFreqCmd)
	freqCmd.AddCommand(sforlFreqCmd)
}

func init() {
	initEuroCmd()
	initSForLCmd()
	initMatchCmd()
	initFreqCmd()
	rootCmd.AddCommand(freqCmd)
	rootCmd.AddCommand(matchCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
