package main

import "fmt"

func main() {
	//test every func here
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr", i)

	fmt.Println("pointer:", &i)

}
