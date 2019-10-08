package main

import (
	"flag"
	"github.com/anexia-it/golang-intro19/intro19"
	"log"
)

var serverAddr = "127.0.0.1:9000"

func init() {
	flag.StringVar(&serverAddr, "addr", serverAddr, "configures server listening address")
}

func main() {
	flag.Parse()

	if err := intro19.RunServer(serverAddr); err != nil {
		log.Fatal(err)
	}
}
