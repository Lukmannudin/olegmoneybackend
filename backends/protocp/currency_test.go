package protocp

import (
	"testing"

	fuzz "github.com/google/gofuzz"
	"github.com/luno/jettison/errors"
	"github.com/luno/jettison/jtest"
	"github.com/stretchr/testify/assert"

	"bitx/currency"
)

func TestDecimalE8FromProtoString(t *testing.T) {
	testCases := []struct {
		name   string
		in     string
		expOut currency.DecimalE8
		expErr error
	}{
		{
			name:   "empty string",
			expOut: currency.ZeroD8,
		},
		{
			name:   "string with letters",
			in:     "12j34",
			expErr: errors.New("invalid string"),
		},
		{
			name:   "string with invalid chars",
			in:     "1,3",
			expErr: errors.New("invalid string"),
		},
		{
			name:   "string with a big decimal number",
			in:     "112233.98765432101",
			expErr: currency.ErrPrecisionLoss,
		},
		{
			name:   "string with a decimal number",
			in:     "112233.9876",
			expOut: currency.FToD8(112233.9876),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualOut, actualErr := DecimalE8FromProtoString(tc.in)
			if tc.expErr != nil {
				assert.Error(t, actualErr)
				jtest.Assert(t, tc.expErr, actualErr)
			} else {
				assert.NoError(t, actualErr)
				assert.Equal(t, tc.expOut, actualOut)
			}
		})
	}
}

func TestBigDecimalFromProto(t *testing.T) {
	testCases := []struct {
		name   string
		in     string
		expOut currency.BigDecimal
		expErr error
	}{
		{
			name:   "empty string",
			expOut: currency.Zero(),
		},
		{
			name:   "string with letters",
			in:     "12j34",
			expErr: errors.New("invalid string"),
		},
		{
			name:   "string with invalid chars",
			in:     "1,3",
			expErr: errors.New("invalid string"),
		},
		{
			name:   "string with a big decimal number",
			in:     "112233.98765432101",
			expOut: currency.NewFromFloat64(112233.98765432101, 11),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualOut, actualErr := BigDecimalFromProto(tc.in)
			if tc.expErr != nil {
				assert.Error(t, actualErr)
				jtest.Assert(t, tc.expErr, actualErr)
			} else {
				assert.NoError(t, actualErr)
				assert.Equal(t, tc.expOut, actualOut)
			}
		})
	}
}

func BenchmarkSerialisation(b *testing.B) {
	bdFuzz := fuzz.New().Funcs(
		func(bd *currency.BigDecimal, c fuzz.Continue) {
			*bd = currency.NewBigDecimalFromInt64(c.Int63(), currency.Scale(c.Intn(18)))
			if c.RandBool() {
				*bd = bd.Neg()
			}
		},
	)

	benchCases := []struct {
		name string
		f    func(*testing.B, currency.BigDecimal)
	}{
		{name: "string", f: func(b *testing.B, d currency.BigDecimal) {
			s, _ := BigDecimalToProto(d)
			_, _ = BigDecimalFromProto(s)
		}},
		{name: "ints", f: func(b *testing.B, d currency.BigDecimal) {
			pb, _ := BigDecimalToDecimalProto(d)
			_, _ = BigDecimalFromDecimalProto(pb)
		}},
	}

	for _, bc := range benchCases {
		b.Run(bc.name, func(b *testing.B) {
			var bd currency.BigDecimal
			bdFuzz.Fuzz(&bd)
			for i := 0; i < b.N; i++ {
				bc.f(b, bd)
			}
		})
	}
}
