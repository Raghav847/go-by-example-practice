package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type QuoteData struct {
	ID     int    `json:"id"`
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

type QuoteResponse struct {
	Quotes []QuoteData `json:"quotes"`
}

func getData(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://dummyjson.com/quotes")
	if err != nil {
		http.Error(w, "cannot get data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "cannot read data", http.StatusInternalServerError)
		return
	}

	var prettyResp bytes.Buffer
	err = json.Indent(&prettyResp, body, "", "  ")
	if err != nil {
		http.Error(w, "cannot make it pretty", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(prettyResp.Bytes())
}

func main() {
	http.HandleFunc("/", getData)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
