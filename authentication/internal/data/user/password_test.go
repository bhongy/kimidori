package user_test

import (
	"testing"

	"github.com/bhongy/kimidori/authentication/internal/data/user"
)

func TestNewPassword(t *testing.T) {
	plainText := "test-password-7oDGpy8iv"

	p1, _ := user.NewPassword(plainText)
	p2, _ := user.NewPassword(plainText)

	if p1.String() == plainText {
		t.Error("Expect the password to be hashed")
	}

	if p1.String() == p2.String() {
		t.Errorf("Expect different hashes: %q", p1.String())
	}
}

func TestPassword_Compare(t *testing.T) {
	plainText := "test-password-7oDGpy8iv"
	p, _ := user.NewPassword(plainText)

	t.Run("match", func(t *testing.T) {
		t.Parallel()

		if !p.Compare(plainText) {
			t.Errorf("Expect the same password to match")
		}
	}) // t.Run("match", ...)

	t.Run("no match", func(t *testing.T) {
		t.Parallel()

		if p.Compare("foobar") {
			t.Errorf("Expect different passwords not to match")
		}
	}) // t.Run("no match", ...)
}
