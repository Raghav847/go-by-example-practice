package main

import "fmt"

func rangeTypes() int {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	fmt.Println("sum:", rangeTypes())

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "go" {
		fmt.Println(i, string(c))
	}
}
