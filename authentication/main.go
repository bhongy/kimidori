package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

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
	mux.Handle("/user", createUser{userService})

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

type createUser struct {
	userService user.Service
}

func (h createUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	u, err := h.userService.Signup("bhongy", "samui.seadog")
	if err != nil {
		log.Println("Error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := json.Marshal(struct {
		ID        string    `json:"id"`
		Username  string    `json:"username"`
		CreatedAt time.Time `json:"createdAt"`
	}{
		u.ID, u.Username, u.CreatedAt,
	})

	if err != nil {
		log.Println("Error: encode body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	header := w.Header()
	header.Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}
