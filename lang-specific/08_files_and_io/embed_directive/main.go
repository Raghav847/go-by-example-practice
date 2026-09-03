package main

import (
	_ "embed"
	"fmt"
)

//go:embed assets/hello.txt
var hello string

func main() {
	fmt.Print(hello)
}
