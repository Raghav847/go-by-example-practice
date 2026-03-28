package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func console(str string) {
	fmt.Println(str)
	return
}


func functions() { 
	res := plus(2, 3)
	fmt.Println("1+2 =", res)

	res2 := plusPlus(1, 2, 3)
	fmt.Println(res2)
	str := "hehe"
	console(str)
}
