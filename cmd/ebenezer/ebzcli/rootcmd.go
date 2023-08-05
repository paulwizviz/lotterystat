package main

import (
	"log"

	"paulwizviz/lotterystat/internal/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   appName,
	Short: "Ebenezer cli is a command line app to analyse UK lottery results",
	Long: `A command line app for users to extract csv data from UK lottery site,
store and analyse results locally.`,
}

func rootCmdIni() {
	rootCmd.AddCommand(euroCmd)
}

func configViper() {
	p, err := config.Path()
	if err != nil {
		log.Fatal(err)
	}

	viper.AddConfigPath(p)
	viper.SetConfigName(config.SettingFileName)
	viper.SetConfigType(config.SettingFileType)
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(configViper)

	rootCmdIni()
}
