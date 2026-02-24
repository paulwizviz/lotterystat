package ebzcli

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/paulwizviz/lotterystat/internal/ebzrest"
	"github.com/paulwizviz/lotterystat/internal/ebzweb"
	"github.com/paulwizviz/lotterystat/internal/sqlops"
	"github.com/spf13/viper"
)

func runWebserver(port int) {
	db, err := sqlops.NewSQLiteFile(viper.GetString("database_path"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mux := http.NewServeMux()
	mux = ebzrest.New(mux, db)
	mux = ebzweb.New(mux)
	fmt.Printf("Listening on port: %d", port)
	err = http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), mux)
	if err != nil {
		log.Fatal(err)
	}
}

func availablePort() int {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	addr := listener.Addr().(*net.TCPAddr)
	return addr.Port
}
