package main

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"fmt"
	"io"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	people := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Cindy", Age: 15},
	}
	out := new(strings.Builder)
	enc := jsontext.NewEncoder(out)

	for _, p := range people {
		json.MarshalEncode(enc, p)
	}
	fmt.Println(out.String())

	in := strings.NewReader(`
    {"Name":"Alice","Age":25}
    {"Name":"Bob","Age":30}
    {"Name":"Cindy","Age":15}
	`)
	dec := jsontext.NewDecoder(in)
	for {
		var p Person
		err := json.UnmarshalDecode(dec, &p)
		if err == io.EOF {
			break
		}
		fmt.Println(p)
	}
}
