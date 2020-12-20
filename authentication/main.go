package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	addr := "localhost:8081"
	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	log.Println(fmt.Sprintf("Server is running at: %s", addr))
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from authentication!"))
}

func createAccount(w http.ResponseWriter, r *http.Request) {

}
