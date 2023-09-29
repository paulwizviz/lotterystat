package main

import (
	"log"
	"paulwizviz/lotterystat/internal/config"
)

func main() {
	err := config.Initialize()
	if err != nil {
		log.Fatal(err)
	}
}
