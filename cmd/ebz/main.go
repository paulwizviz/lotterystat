package main

import (
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

func init() {
	initDBCmd()
	persistDBCmd()
	rootCmd.AddCommand(dbCmd)
}

func execute() error {
	return rootCmd.Execute()
}

func main() {
	log.Println("Hello")

	execute()

	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt)
	// <-c
}
