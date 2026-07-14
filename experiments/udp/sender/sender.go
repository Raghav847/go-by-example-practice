package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	message := []byte("Hello There!!")
	_, err = conn.Write(message)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Message sent!")
}
