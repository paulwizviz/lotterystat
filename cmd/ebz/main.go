package main

import (
	"log"
	"net/http"

	"github.com/paulwizviz/lotterystat/internal/ebzmux"
)

func main() {
	mux := ebzmux.New()
	err := http.ListenAndServe("0.0.0.0:8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
