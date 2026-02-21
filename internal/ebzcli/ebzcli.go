package ebzcli

import (
	"fmt"
	"log"

	"github.com/paulwizviz/lotterystat/internal/ebzconfig"
	"github.com/spf13/cobra"
)

var start bool

func init() {
	err := ebzconfig.Initialize()
	if err != nil {
		log.Println(err)
	}
	rootCmd.Flags().BoolVarP(&start, "start", "s", false, "Start the frontend web server")
}

var rootCmd = &cobra.Command{
	Use:   "ebz",
	Short: "ebz is a cli app to help you analyze UK National Lottery results.",
	Run: func(cmd *cobra.Command, args []string) {
		if start {
			port := availablePort()
			rawUrl := fmt.Sprintf("http://localhost:%d", port)
			openBrowser(rawUrl)
			runWebserver(port)
			return
		}
		cmd.Help()
	},
}

func Execute() error {
	rootCmd.AddCommand(tballCmd)
	rootCmd.AddCommand(euroCmd)
	rootCmd.AddCommand(lottoCmd)
	rootCmd.AddCommand(sflifeCmd)
	return rootCmd.Execute()
}
