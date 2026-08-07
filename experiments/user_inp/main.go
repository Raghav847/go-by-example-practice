package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readScanner() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your alias: ")
	scanner.Scan()
	name := scanner.Text()
	fmt.Println("Hello, ", name)
}

func readLine() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter something: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	fmt.Println("You said:", input)
}

func main() {
	//readScanner()
	readLine()
}
