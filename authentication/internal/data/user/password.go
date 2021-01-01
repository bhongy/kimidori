package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// password is the hash of plain text password
type password struct {
	hashed string
}

// NewPassword creates a new hash of a plain-text password
func NewPassword(plainText string) (password, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)
	if err != nil {
		return password{}, fmt.Errorf("new password: bcrypt.GenerateFromPassword error: %v", err)
	}
	p := password{string(h)}
	return p, nil
}

// String returns a string representation of the Password
// intended for storing in the database
//
// Do not use it to compare passwords
func (p password) String() string {
	return p.hashed
}

// CheckPassword checks the given plain-text password against the hashed password
func CheckPassword(hashed, plainText string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashed),
		[]byte(plainText),
	)
	return err == nil
}
