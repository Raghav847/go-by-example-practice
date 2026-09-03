package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/posts"

type data struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var fullData []data

func getData() ([]data, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error getting data")
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&fullData)
	if err != nil {
		fmt.Println("error decoding data")
		return nil, err
	}
	return fullData, nil
}

func handle(w http.ResponseWriter, r *http.Request) {
	resp, err := getData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := json.MarshalIndent(resp, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func main() {
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}
