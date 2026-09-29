package main

import (
	"errors"
	"fmt"
)

func main() {
	s := NewStore()
	s.Set("a", "42")
	s.Set("b", "72")
	s.Delete("a")

	val, err := s.Get("b")
	if err != nil {
		if errors.Is(err, ErrKeyDoesNotExist) {
			return
		}
		fmt.Println(err)
	}
	fmt.Println(val)
}
