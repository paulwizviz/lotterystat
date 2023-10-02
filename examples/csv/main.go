package main

import (
	"context"
	"fmt"
	"log"
	"paulwizviz/lotterystat/internal/csvutil"
	"paulwizviz/lotterystat/internal/euro"
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
}
