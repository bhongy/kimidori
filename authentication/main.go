package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bhongy/kimidori/authentication/internal/data/postgres"
	"github.com/bhongy/kimidori/authentication/internal/data/user"
	"github.com/google/uuid"
)

func main() {
	db, err := postgres.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	ur := user.NewRepository(db)
	mux.Handle("/user", createUser{ur})

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
	userRepo user.Repository
}

func (h createUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
		return
	}

	password, err := user.NewPassword("samui.seadog")
	if err != nil {
		log.Printf("Error: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u := user.User{
		UUID:      uuid.New(),
		Username:  "bhongy",
		Password:  password,
		CreatedAt: time.Now(),
	}
	if err = h.userRepo.Create(&u); err != nil {
		log.Printf("Error: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := json.Marshal(struct {
		ID        int       `json:"id"`
		UUID      uuid.UUID `json:"uuid"`
		Username  string    `json:"username"`
		CreatedAt time.Time `json:"createdAt"`
	}{
		u.ID, u.UUID, u.Username, u.CreatedAt,
	})

	if err != nil {
		log.Printf("Error: encode body: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	header := w.Header()
	location := fmt.Sprintf("/user/%d", u.ID)
	header.Set("Location", location)
	header.Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}
