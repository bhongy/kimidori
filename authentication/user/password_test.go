package user_test

import (
	"testing"

	"github.com/bhongy/kimidori/authentication/user"
)

func TestNewPassword(t *testing.T) {
	plainText := "test-password-7oDGpy8iv"

	p1, _ := user.NewPassword(plainText)
	p2, _ := user.NewPassword(plainText)

	// outside this test, do not compare passwords using `.String()`
	if p1.String() == plainText {
		t.Error("expect the password to be hashed")
	}

	if p1.String() == p2.String() {
		t.Errorf("expect different hashes: %q", p1.String())
	}
}

func TestCheckPassword(t *testing.T) {
	plainText := "test-password-7oDGpy8iv"
	hashed, _ := user.NewPassword(plainText)

	t.Run("match", func(t *testing.T) {
		t.Parallel()
		if !user.CheckPassword(hashed, plainText) {
			t.Errorf("Expect the same password to match")
		}
	})

	t.Run("no match", func(t *testing.T) {
		t.Parallel()
		if user.CheckPassword(hashed, "foobar") {
			t.Errorf("Expect different passwords not to match")
		}
	})
}

func TestPassword_Check(t *testing.T) {
	plainText := "test-password-7oDGpy8iv"
	p, _ := user.NewPassword(plainText)

	t.Run("match", func(t *testing.T) {
		t.Parallel()
		if !p.Check(plainText) {
			t.Errorf("Expect the same password to match")
		}
	})

	t.Run("no match", func(t *testing.T) {
		t.Parallel()
		if p.Check("foobar") {
			t.Errorf("Expect different passwords not to match")
		}
	})
}
