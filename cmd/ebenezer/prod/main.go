package main

import (
	"log"
	"paulwizviz/lotterystat/internal/config"
)

func main() {
	err := config.Initilalize()
	if err != nil {
		log.Fatal(err)
	}
	Execute()
}
