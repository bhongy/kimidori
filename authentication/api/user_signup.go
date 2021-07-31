package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/bhongy/kimidori/authentication/user"
)

func UserSignup(userService user.Service) http.Handler {
	return &userSignupHandler{userService}
}

type UserSignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserSignupResponse struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	CreatedAt string `json:"createdAt"`
}

type userSignupHandler struct {
	userService user.Service
}

func (h *userSignupHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		code := http.StatusMethodNotAllowed
		http.Error(w, http.StatusText(code), code)
		return
	}

	var req UserSignupRequest
	err := decodeJSONBody(r, &req)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.Error(), http.StatusBadRequest)
			return
		}
		log.Println("decode request body:", err)
		code := http.StatusInternalServerError
		http.Error(w, http.StatusText(code), code)
		return
	}

	u, err := h.userService.Signup(req.Username, req.Password)
	if err != nil {
		if err == user.ErrUsernameExists {
			http.Error(w, "username already exists", http.StatusBadRequest)
			return
		}
		log.Println("user service signup:", err)
		code := http.StatusInternalServerError
		http.Error(w, http.StatusText(code), code)
		return
	}

	body, err := json.Marshal(UserSignupResponse{
		ID:        u.ID,
		Username:  u.Username,
		CreatedAt: u.CreatedAt.String(),
	})

	if err != nil {
		// log server-side too because we don't expect that this could happen
		log.Println("marshal User to body:", err)
		http.Error(w, "username already exists", http.StatusBadRequest)
		return
	}

	header := w.Header()
	header.Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}
