// Package ebzcli contains implementation of cli operation
package ebzcli

import (
	"log"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	viperConfigPathKey = "config-path"
)

const (
	configDirName = ".ebz"
	appName       = "ebzcli"
)

var (
	// location where the state and local db file is stored
	configPath string
)

var (
	euroCmd = &cobra.Command{
		Use:   "euro",
		Short: "Sub command to analyse Euro Draw Results",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
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
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	configPath = path.Join(homeDir, configDirName)
	_, err = os.Stat(configPath)
	if os.IsNotExist(err) {
		err := os.Mkdir(configPath, 0775)
		if err != nil {
			log.Fatal(err)
		}
	}
	viper.Set(viperConfigPathKey, configPath)

}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(euroCmd)

}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
