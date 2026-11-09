package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const apiKey = ""

// movie struct and json tags inspect output
type movie struct {
	Title  string `json:"Title"`
	Poster string `json:"Poster"`
}

func Poster(posterURL string) error {
	resp, err := http.Get(posterURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("poster download failed: %s", resp.Status)
	}

	file, err := os.Create("poster.jpg")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("Poster downloaded!!")
	return nil
}

func main() {
	//var URL = "http://img.omdbapi.com/?apikey=[yourkey]&"
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a movie name: ")
	movieName, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
		return
	}
	movieName = strings.TrimSpace(movieName)
	fmt.Println("Enter a movie year: ")
	movieYear, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
		return
	}
	movieYear = strings.TrimSpace(movieYear)

	params := url.Values{}
	params.Set("t", movieName)
	params.Set("y", movieYear)
	params.Set("apikey", apiKey)

	apiURL := "https://www.omdbapi.com/?" + params.Encode()
	resp, err := http.Get(apiURL)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()
	var m movie
	err = json.NewDecoder(resp.Body).Decode(&m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m.Title)
	fmt.Println(m.Poster)
	err = Poster(m.Poster)
}
