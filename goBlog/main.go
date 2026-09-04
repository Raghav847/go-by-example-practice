package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/yuin/goldmark"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /posts/{slug}", PostHandler(FileReader{}))

	addr := flag.String("addr", ":3030", "server address")
	flag.Parse()

	log.Printf("Server Started on port: %v", *addr)
	log.Fatal(http.ListenAndServe(*addr, mux))

}

type SlugReader interface {
	Read(slug string) (string, error)
}

type FileReader struct{}

func (fr FileReader) Read(slug string) (string, error) {
	f, err := os.Open(slug + ".md")
	if err != nil {
		return "", err
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func PostHandler(sl SlugReader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug := r.PathValue("slug")
		postMarkdown, err := sl.Read(slug)
		if err != nil {
			http.Error(w, "Post not found", http.StatusNotFound)
			return
		}
		var buf bytes.Buffer
		err = goldmark.Convert([]byte(postMarkdown), &buf)
		if err != nil {
			panic(err)
		}
		io.Copy(w, &buf)
	}
}
