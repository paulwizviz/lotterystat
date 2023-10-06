package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var euroCmd = &cobra.Command{
	Use:   "euro",
	Short: "euro is a sub command of ebz app to perform analysis",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			fmt.Println(args)
			return
		}
		cmd.Help()
	},
}

var rootCmd = &cobra.Command{
	Use:   "ebz",
	Short: "ebz is a cli app to help you analyze UK National Lottery statistics",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			cmd.Help()
			return
		}
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(euroCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
