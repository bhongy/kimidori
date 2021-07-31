package user_test

import (
	"testing"
	"time"

	"github.com/bhongy/kimidori/authentication/user"
)

const referenceTime = "2020-04-12T08:49:52Z"

func newReferenceTime(t *testing.T) time.Time {
	tt, err := time.Parse(time.RFC3339, referenceTime)
	if err != nil {
		t.Fatalf("Cannot parse time: %q", tt)
	}
	return tt
}

// TODO: test Equal method

func Test_Timestamp_Scan(t *testing.T) {
	ts := user.NewTimestamp(time.Now())
	tt := newReferenceTime(t)
	err := ts.Scan(tt)
	if err != nil {
		t.Error(err)
	}
	if got, want := ts.String(), referenceTime; got != want {
		t.Errorf("Scanned value incorrect. want: %v, got: %v", got, want)
	}
}

func Test_Timestamp_String(t *testing.T) {
	tt := newReferenceTime(t)
	ts := user.NewTimestamp(tt)
	if formatted := ts.String(); formatted != referenceTime {
		t.Error(formatted)
	}
}
