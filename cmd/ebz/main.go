package main

import "github.com/paulwizviz/lotterystat/internal/ebzcli"

func main() {

	ebzcli.Execute()

	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt)
	// <-c
}
