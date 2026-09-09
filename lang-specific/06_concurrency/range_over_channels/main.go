package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "big"
	queue <- "me"
	close(queue)

	for str := range queue {
		fmt.Println(str)
	}
}
