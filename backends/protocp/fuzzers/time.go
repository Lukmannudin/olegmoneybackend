package fuzzers

import (
	"time"

	fuzz "github.com/google/gofuzz"
)

func FuzzTime(t *time.Time, c fuzz.Continue) {
	d := time.Duration(c.Intn(100)) * time.Millisecond
	*t = time.Now().Add(d).UTC()
}

func FuzzDuration(d *time.Duration, c fuzz.Continue) {
	*d = time.Duration(c.Intn(100)) * time.Millisecond
}
