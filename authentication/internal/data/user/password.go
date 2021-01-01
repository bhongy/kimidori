package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// password is the hash of plain text password
type password []byte

// NewPassword creates a new hash of a plain-text password
func NewPassword(plainText string) (password, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)
	if err != nil {
		return password{}, fmt.Errorf("new password: bcrypt.GenerateFromPassword error: %v", err)
	}
	return password(h), nil
}

// String returns a string representation of the Password
// intended for storing in the database
//
// Do not use it to compare passwords
func (p password) String() string {
	return string(p)
}

// CheckPassword checks the given plain-text password against the hashed password
func CheckPassword(hashed, plainText []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashed, plainText)
	return err == nil
}
