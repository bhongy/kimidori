package api_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/bhongy/kimidori/authentication/api"
	"github.com/bhongy/kimidori/authentication/user"
)

type mockUserService struct {
	createdUser user.User
}

func (svc *mockUserService) Signup(username, password string) (user.User, error) {
	return svc.createdUser, nil
}

func TestUserSignup(t *testing.T) {
	u := user.User{
		ID:        "fake.id",
		Username:  "fake.username",
		CreatedAt: user.NewTimestamp(time.Now()),
	}

	body, err := json.Marshal(api.UserSignupRequest{
		Username: u.Username,
		Password: "fake.password",
	})
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}

	req, err := http.NewRequest("POST", "/does-not-matter", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("create request: %v", err)
	}

	recorder := httptest.NewRecorder()
	userService := mockUserService{createdUser: u}
	h := api.UserSignup(&userService)

	h.ServeHTTP(recorder, req)

	expectedStatus := http.StatusCreated
	if status := recorder.Code; status != expectedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, expectedStatus)
	}

	expectedBody := fmt.Sprintf(
		`{"id":"%v","username":"%v","createdAt":"%v"}`,
		u.ID,
		u.Username,
		u.CreatedAt.String(),
	)
	if body := recorder.Body.String(); body != expectedBody {
		t.Errorf("handler returned unexpected body: got %v want %v",
			body, expectedBody)
	}
}
