package main

import "net/http"

func Serve(pngData []byte) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/mandelbrot.png", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		w.Header().Set("Cache-Control", "public, max-age=3600")

		if _, err := w.Write(pngData); err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
	})
	return http.ListenAndServe(":8000", mux)
}
