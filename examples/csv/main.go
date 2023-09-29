package main

import (
	"context"
	"fmt"
	"log"
	"paulwizviz/lotterystat/internal/csvproc"
	"paulwizviz/lotterystat/internal/euro"
)

func main() {
	r, err := csvproc.DownloadFrom(euro.CSVUrl)
	if err != nil {
		log.Fatal(err)
	}
	ecd := csvproc.EuroCSV(context.TODO(), r)
	for c := range ecd {
		fmt.Println(c)
	}
}
