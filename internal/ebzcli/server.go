package ebzcli

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/paulwizviz/lotterystat/internal/ebzweb"
)

func runWebserver(port int) {
	mux := http.NewServeMux()
	mux = ebzweb.New(mux)
	fmt.Printf("Listening on port: %d", port)
	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), mux)
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
