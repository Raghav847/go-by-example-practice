package main

import (
	"fmt"
	"math"
)

func Constants() {
	const s string = "constant"
	fmt.Println(s)

	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(math.Sin(n))
}
