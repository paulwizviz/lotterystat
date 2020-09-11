package ebzcli

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "ebz",
	Short: "ebz is a cli app to help you analyze UK National Lottery results.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() error {
	return rootCmd.Execute()
}
