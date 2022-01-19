package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/bhongy/kimidori/go-api/handlers"
	"github.com/gorilla/mux"
)

var shutdownTimeout time.Duration

func init() {
	flag.DurationVar(
		&shutdownTimeout, "graceful-timeout", 30*time.Second,
		"the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
}

func main() {
	flag.Parse()

	l := log.New(os.Stdout, "[product-api] ", log.LstdFlags)
	r := mux.NewRouter()

	pr := r.PathPrefix("/products").Subrouter()
	{
		ph := handlers.NewProducts(l)
		pr.Methods("GET").HandlerFunc(ph.Get)
		pr.Methods("POST").HandlerFunc(ph.Post)
		pr.Path("/{id:[0-9]+}").HandlerFunc(ph.Put)
	}

	s := http.Server{
		Addr:         "127.0.0.1:8080",
		Handler:      r,
		ErrorLog:     l,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		l.Println("Starting server on port 8080")

		if err := s.ListenAndServe(); err != nil {
			l.Fatalf("Error starting server: %s\n", err)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	sig := <-sigCh
	l.Println("Received terminate, graceful shutdown", sig)

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	s.Shutdown(ctx)
}
