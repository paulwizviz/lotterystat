package ebzcli

import (
	"fmt"
	"log"

	"github.com/paulwizviz/lotterystat/internal/ebzconfig"
	"github.com/spf13/cobra"
)

func init() {
	err := ebzconfig.Initialize()
	if err != nil {
		log.Println(err)
	}
}

var rootCmd = &cobra.Command{
	Use:   "ebz",
	Short: "ebz is a cli app to help you analyze UK National Lottery results.",
	Run: func(cmd *cobra.Command, args []string) {
		port := availablePort()
		rawUrl := fmt.Sprintf("http://localhost:%d", port)
		openBrowser(rawUrl)
		runWebserver(port)
	},
}

var tballCmd = &cobra.Command{
	Use:   "tball",
	Short: "tball is a subcommand",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tball")
	},
}

func Execute() error {
	rootCmd.AddCommand(tballCmd)
	return rootCmd.Execute()
}
