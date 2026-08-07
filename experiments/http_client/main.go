package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const apiKey string = "" //tmdb access token

type PopularMovieResponse struct {
	Page    int     `json:"page"`
	Results []Movie `json:"results"`
}

type Movie struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	PosterPath  string  `json:"poster_path"`
	ReleaseDate string  `json:"release_date"`
	Popularity  float64 `json:"popularity"`
	VoteAverage float64 `json:"vote_average"`
}

func main() {
	url := "https://api.themoviedb.org/3/movie/popular"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("err in fetching %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("err fetch status %v", resp.StatusCode)
		return
	}

	var result PopularMovieResponse
	json.NewDecoder(resp.Body).Decode(&result)
	pretty, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		fmt.Println("cant pretty")
		return
	}
	fmt.Println(string(pretty))
}
