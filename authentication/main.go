package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bhongy/kimidori/authentication/api"
	"github.com/bhongy/kimidori/authentication/repository/postgres"
	"github.com/bhongy/kimidori/authentication/repository/postgres/db"
	"github.com/bhongy/kimidori/authentication/user"
)

func main() {
	conn, err := db.Open()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	userRepo := postgres.NewUserRepository(conn)
	userService := user.NewService(userRepo)
	mux.Handle("/user/signup", api.UserSignup(userService))

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
