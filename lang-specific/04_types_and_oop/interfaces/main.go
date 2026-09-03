package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
)

type Customer struct {
	Name string
	Age  int
}

func (c *Customer) WriteJSON(w io.Writer) error {
	js, err := json.Marshal(c)
	if err != nil {
		return err
	}

	_, err = w.Write(js)
	return err
}

func main() {
	c := &Customer{Name: "Bane", Age: 25}

	var buf bytes.Buffer
	err := c.WriteJSON(&buf)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("/tmp/customer")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = c.WriteJSON(f)
	if err != nil {
		log.Fatal(err)
	}
}
