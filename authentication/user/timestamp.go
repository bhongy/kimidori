package user

import (
	"fmt"
	"time"
)

// we want to consistently format the time when serializing it
// e.g. when saving it to DB or sending it over HTTP

type timestamp time.Time

func NewTimestamp(t time.Time) timestamp {
	return timestamp(t)
}

// Equal allows tools like `cmp` to compare two timestamps for equality
func (t timestamp) Equal(other timestamp) bool {
	return t.String() == other.String()
}

// Scan implements Scanner interface supporting types: time.Time
func (t *timestamp) Scan(v interface{}) error {
	if v == nil {
		return nil
	}
	if v, ok := v.(time.Time); ok {
		*t = NewTimestamp(v)
		return nil
	}
	return fmt.Errorf("Can not scan %v to timestamp", v)
}

func (t timestamp) String() string {
	return time.Time(t).Format(time.RFC3339)
}
