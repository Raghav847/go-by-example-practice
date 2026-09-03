package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี"

	fmt.Println("Len: ", len(s))

	for i, _ := range s {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runeVal := range s {
		fmt.Printf("%#U starts at %d\n", runeVal, idx)
	}
}
