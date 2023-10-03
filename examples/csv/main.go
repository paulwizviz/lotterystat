package main

import (
	"context"
	"fmt"
	"log"
	"paulwizviz/lotterystat/internal/csvutil"
	"paulwizviz/lotterystat/internal/euro"
	"paulwizviz/lotterystat/internal/sforl"
)

func main() {
	r, err := csvutil.DownloadFrom(euro.CSVUrl)
	if err != nil {
		log.Fatal(err)
	}
	ecd := euro.ProcessCSV(context.TODO(), r)
	for c := range ecd {
		fmt.Println(c)
	}
	fmt.Println("--------------")
	r, err = csvutil.DownloadFrom(sforl.CSVUrl)
	if err != nil {
		log.Fatal(err)
	}
	scd := sforl.ProcessCSV(context.TODO(), r)
	for c := range scd {
		fmt.Println(c)
	}
}
