package main

import (
	"log"

	"github.com/pasca-l/global-ip-retriever/server"
)

func main() {
	err := server.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
