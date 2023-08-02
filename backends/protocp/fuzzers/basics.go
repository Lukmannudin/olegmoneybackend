package fuzzers

import fuzz "github.com/google/gofuzz"

func FuzzInt(i *int, c fuzz.Continue) {
	*i = c.Intn(1000)
}
