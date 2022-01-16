package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/bhongy/kimidori/go-api/handlers"
)

func main() {
	l := log.New(os.Stdout, "[product-api] ", log.LstdFlags)
	mux := http.NewServeMux()
	mux.Handle("/", handlers.NewProducts(l))

	s := http.Server{
		Addr:         "127.0.0.1:8080",
		Handler:      mux,
		ErrorLog:     l,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		l.Println("Starting server on port 8080")

		err := s.ListenAndServe()
		if err != nil {
			l.Fatalf("Error starting server: %s\n", err)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	sig := <-sigCh
	l.Println("Received terminate, graceful shutdown", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
