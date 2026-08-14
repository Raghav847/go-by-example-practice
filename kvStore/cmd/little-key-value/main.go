package main

import (
	"fmt"
	"log"

	"github.com/Raghav847/go-by-example-practice/kvStore/internal/api"
)

func main() {
	fmt.Println("Starting KVStore...")

	server := api.New()
	err := server.Serve(":8888")
	if err != nil {
		log.Fatal(err)
	}
}
