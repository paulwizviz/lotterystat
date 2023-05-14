package main

import (
	"log"

	"github.com/paulwizviz/lotterystat/internal/config"
)

const (
	appName = "ebzcli"
)

func main() {

	configPath, err := config.Path()
	if err != nil {
		log.Fatal(err)
	}

	err = config.Initialize(configPath)
	if err != nil {
		log.Fatal(err)
	}

	err = rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
