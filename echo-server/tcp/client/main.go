package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("Hello, World!"))
	if err != nil {
		log.Fatal(err)
	}

	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)
	fmt.Println(string(buffer))
}
