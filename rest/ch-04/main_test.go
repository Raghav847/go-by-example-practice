package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddCacheHeaders(t *testing.T) {
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	req := httptest.NewRequest("GET", "/lists", nil)
	rec := httptest.NewRecorder()
	handler := addCacheHeaders(testHandler)
	handler(rec, req)
	if rec.Header().Get("Cache-Control") != "public, max-age=300" {
		t.Errorf("Not valid Cache-Control found, got %v, want %v", rec.Header().Get("Cache-Control"), "public, max-age=300")
	}
	if rec.Header().Get("Expires") == "" {
		t.Errorf("Not valid Expires, got %v, want not empty", rec.Header().Get("Expires"))
	}
}

func TestLoginAPI(t *testing.T) {
	var err error
	repository, err = NewRepository("file::memory:?cache=shared")
	if err != nil {
		t.Errorf("NewRepository() error = %v", err)
	}
	repository.Init()

	req := httptest.NewRequest("POST", "/login", strings.NewReader(`{"username":"admin","password":"admin"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	handleLogin(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("handleLogin() status = %v, want %v", rec.Code, http.StatusOK)
	}
}

func makeRequest(method, url string, body io.Reader, token string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	return http.DefaultClient.Do(req)
}

func TestLoginAndList(t *testing.T) {
	baseURL := "http://localhost:8888"

	reqBody := strings.NewReader(`{"username":"admin","password":"admin"}`)
	resp, err := makeRequest("POST", baseURL+"/login", reqBody, "")
	if err != nil {
		t.Fatalf("Failed to make a request: %v", err)
	}
	defer resp.Body.Close()

	var response map[string]string
	json.NewDecoder(resp.Body).Decode(&response)
	token := response["token"]
	if token == "" {
		t.Fatal("No token returned from login")
	}
}
