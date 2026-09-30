package main

import (
	"errors"
	"fmt"
)

func main() {
	s := NewStore(2)
	_ = s.Set("a", "42")
	_ = s.Set("b", "72")

	if err := s.Set("a", "8"); err != nil {
		fmt.Println(err)
		return
	}

	val, err := s.Get("b")
	if err != nil {
		if errors.Is(err, ErrKeyDoesNotExist) {
			return
		}
		fmt.Println(err)
	}
	fmt.Println(val)
}
