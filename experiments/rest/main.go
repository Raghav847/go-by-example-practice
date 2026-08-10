package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Event struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Genres      []string `json:"genres"`
	Duration    int      `json:"duration"`
	OrganizedBy string   `json:"organized_by"`
}

type APIResponse struct {
	Page       int     `json:"page"`
	PerPage    int     `json:"per_page"`
	Total      int     `json:"total"`
	TotalPages int     `json:"total_pages"`
	Data       []Event `json:"data"`
}

func longestDuration(organizer, genre string) string {
	baseURL := "https://jsonmock.hackerrank.com/api/events"

	bestID := ""
	bestDuration := -1

	page := 1
	totalPages := 1

	for page <= totalPages {
		reqURL := fmt.Sprintf("%s?organized_by=%s&page=%d",
			baseURL, url.QueryEscape(organizer), page)

		resp, err := http.Get(reqURL)
		if err != nil {
			fmt.Fprintln(os.Stderr, "request error:", err)
			return "-1"
		}

		var apiResp APIResponse
		if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
			resp.Body.Close()
			fmt.Fprintln(os.Stderr, "decode error:", err)
			return "-1"
		}
		resp.Body.Close()

		totalPages = apiResp.TotalPages

		for _, event := range apiResp.Data {
			if containsGenre(event.Genres, genre) {
				if event.Duration > bestDuration ||
					(event.Duration == bestDuration && (bestID == "" || event.ID < bestID)) {
					bestDuration = event.Duration
					bestID = event.ID
				}
			}
		}

		page++

	}

	if bestID == "" {
		return "-1"
	}
	return bestID
}

func containsGenre(genres []string, target string) bool {
	for _, g := range genres {
		if g == target {
			return true
		}
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	organizer, _ := reader.ReadString('\n')
	organizer = strings.TrimSpace(organizer)

	genre, _ := reader.ReadString('\n')
	genre = strings.TrimSpace(genre)

	result := longestDuration(organizer, genre)
	fmt.Println(result)
}
