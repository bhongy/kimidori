package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// password is the hash of plain text password
type password []byte

// NewPassword creates a new hash of a plain-text password
func NewPassword(plainText string) (password, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("new password: bcrypt.GenerateFromPassword error: %v", err)
	}
	return password(hash), nil
}

// String returns a string representation of the Password
// intended for storing in the database
//
// Do not use it to compare passwords
func (p password) String() string {
	return string(p)
}

// Compare checks the given plain-text password whether it matches the Password
func (p password) Compare(plainText string) bool {
	err := bcrypt.CompareHashAndPassword(p, []byte(plainText))
	return err == nil
}
