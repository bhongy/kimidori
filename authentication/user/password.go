package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// password is the hash of a plain text password.
type password string

// NewPassword hashes plainText to create password.
func NewPassword(plainText string) (password, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("generate hash from plainText: %v", err)
	}
	return password(h), nil
}

// String returns a string representation of p
// intended for storing in the database.
//
// Do not use it to compare passwords
func (p password) String() string {
	return string(p)
}

// Check checks plainText against p.
func (p password) Check(plainText string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(p),
		[]byte(plainText))
	return err == nil
}

// CheckPassword checks whether hashed matches the result of hashing plainText.
func CheckPassword(hashed password, plainText string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashed),
		[]byte(plainText))
	return err == nil
}
