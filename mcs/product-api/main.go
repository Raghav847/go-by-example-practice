package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Raghav847/go-by-example-practice/mcs/product-api/data"
	"github.com/Raghav847/go-by-example-practice/mcs/product-api/handlers"
)

var bindAddress = flag.String("bind-address", ":9090", "Bind address for the server")

func main() {
	flag.Parse()

	l := log.New(os.Stdout, "[products-api]", log.LstdFlags)
	v := data.NewValidation()

	ph := handlers.NewProducts(l, v)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /products", ph.ListAll)
	mux.Handle(
		"POST /products",
		ph.MiddlewareValidateProduct(http.HandlerFunc(ph.Create)),
	)
	mux.Handle(
		"PUT /products",
		ph.MiddlewareValidateProduct(http.HandlerFunc(ph.Update)),
	)

	s := http.Server{
		Addr:         *bindAddress,
		Handler:      mux,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		l.Println("Starting server on port 9090")

		l.Fatal(s.ListenAndServe())
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Println("Got signal", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)

}
