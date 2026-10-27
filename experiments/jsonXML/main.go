package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type Ingredient struct {
	Name   string
	Amount int
}

type Money struct {
	Dollars int `json:"dollars"`
	Cents   int `json:"cents"`
}

type Cake struct {
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
	Description string       `json:"description"`
	Price       Money        `json:"price"`
}

func main() {
	choco := Cake{
		Name: "Chocolate",
		Ingredients: []Ingredient{
			Ingredient{"Chocolate", 1},
			Ingredient{"Butter", 2},
			Ingredient{"Flour", 2},
			Ingredient{"Eggs", 5},
		},
		Description: "A bad chocolate cake",
		Price:       Money{30, 99},
	}

	// marshal to json
	jsonData, err := json.MarshalIndent(choco, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	// write data to file
	err = os.WriteFile("output.json", jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// get the json data
	againData, err := os.ReadFile("output.json")
	if err != nil {
		log.Fatal(err)
	}

	if !json.Valid(againData) {
		log.Fatal("invalid json data")
	}

	var newCake Cake
	err = json.Unmarshal(againData, &newCake)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("=========JSON Cake=========")
	fmt.Println(newCake)

	// xml
	xmlOut, err := xml.MarshalIndent(choco, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("=========XML Cake=========")
	fmt.Println(string(xmlOut))
}
