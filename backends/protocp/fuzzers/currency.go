package fuzzers

import (
	fuzz "github.com/google/gofuzz"

	"bitx/currency"
)

func FuzzBigDecimal(bd *currency.BigDecimal, c fuzz.Continue) {
	*bd = currency.NewBigDecimalFromInt64(c.Int63n(1000), currency.ScaleInteger)
}

func FuzzDecimalE8(e8 *currency.DecimalE8, c fuzz.Continue) {
	*e8 = currency.FToD8(c.Float64())
}

func FuzzScale(s *currency.Scale, c fuzz.Continue) {
	*s = currency.ScaleInteger
}
