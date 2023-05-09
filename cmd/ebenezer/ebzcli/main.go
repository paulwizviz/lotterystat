package main

import (
	"log"

	"github.com/paulwizviz/lotterystat/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	appName = "ebzcli"
)

var (
	euroCmd = &cobra.Command{
		Use:   "euro",
		Short: "A subcommand to process Euro draws",
		Run: func(cmd *cobra.Command, args []string) {
			var settings config.Setting
			err := viper.Unmarshal(&settings)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(settings)
		},
	}

	rootCmd = &cobra.Command{
		Use:   appName,
		Short: "Ebenezer cli is a command line app to analyse UK lottery results",
		Long: `A command line app for users to extract csv data from UK lottery site,
store and analyse results locally.`,
	}
)

func initConfig() {
	err := config.CreateIfNotExist()
	if err != nil {
		log.Fatal(err)
	}

	p, err := config.Path()
	if err != nil {
		log.Fatal(err)
	}

	viper.AddConfigPath(p)
	viper.SetConfigName("settings")
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(euroCmd)
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
