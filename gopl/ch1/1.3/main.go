package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo1() string {
	start := time.Now()
	defer func() {
		fmt.Printf("echo1: %v ns\n", time.Since(start).Nanoseconds())
	}()
	s, sep := "", ""
	for i := 1; i < len(os.Args); i++ {
		s += os.Args[i] + sep
		sep = " "
	}
	return s
}

func echo2() string {
	start := time.Now()
	defer func() {
		fmt.Printf("echo1: %v ns\n", time.Since(start).Nanoseconds())
	}()

	s := strings.Join(os.Args[1:], " ")
	return s
}

func main() {
	echo1()
	echo2()
}
