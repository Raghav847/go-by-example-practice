package main

import "fmt"

func loop() {
	i := 1
	//single condition
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//classical
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	//range
	for i := range 3 {
		fmt.Println("range", i)
	}

	//break
	for {
		fmt.Println("loop")
		break
	}

	//continue
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
