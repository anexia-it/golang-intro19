package main

import (
	"github.com/anexia-it/golang-intro19/intro19"
	"log"
)

func main() {
	if err := intro19.RunServer("127.0.0.1:90000"); err != nil {
		log.Fatal(err)
	}
}
