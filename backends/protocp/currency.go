package protocp

import (
	"bitx/currency"
	"bitx/currency/decimalpb"
)

func DecimalE8ToProto(d currency.DecimalE8) (int64, error) {
	return d.E8Value(), nil
}

func DecimalE8FromProto(e8 int64) (currency.DecimalE8, error) {
	return currency.FromE8Value(e8), nil
}

func DecimalE8ToDecimalProto(d currency.DecimalE8) (*decimalpb.Decimal, error) {
	return d.BigDecimal().Proto(), nil
}

func DecimalE8FromDecimalProto(d *decimalpb.Decimal) (currency.DecimalE8, error) {
	return currency.NewFromProto(d).DecimalE8(), nil
}

func DecimalE8ToProtoString(d currency.DecimalE8) (string, error) {
	return d.String(), nil
}

func DecimalE8FromProtoString(s string) (currency.DecimalE8, error) {
	if s == "" {
		return currency.ZeroD8, nil
	}
	return currency.ParseDecimalE8(s)
}

// BigDecimalToProto converts BigDecimal to string
// Deprecated: Don't use string for decimals,
// rather use bitx/lib/protobuf/ptypes/decimal
func BigDecimalToProto(d currency.BigDecimal) (string, error) {
	return d.String(), nil
}

// BigDecimalFromProto converts string to BigDecimal
// Deprecated: Don't use string for decimals,
// rather use decimalpb.Decimal
func BigDecimalFromProto(d string) (currency.BigDecimal, error) {
	if d == "" {
		return currency.Zero(), nil
	}
	return currency.Parse(d)
}

func BigDecimalToDecimalProto(d currency.BigDecimal) (*decimalpb.Decimal, error) {
	return d.Proto(), nil
}

func BigDecimalFromDecimalProto(d *decimalpb.Decimal) (currency.BigDecimal, error) {
	return currency.NewFromProto(d), nil
}

func DecimalFieldToProto(d currency.DecimalField) (int64, error) {
	return d.DecimalE8().E8Value(), nil
}

func DecimalFieldFromProto(e8 int64) (currency.DecimalField, error) {
	return currency.NewDecimalField(currency.FromE8Value(e8)), nil
}

func NullDecimalFieldToProto(d currency.NullDecimalField) (int64, error) {
	return d.DecimalField.DecimalE8().E8Value(), nil
}

func NullDecimalFieldFromProto(e8 int64) (currency.NullDecimalField, error) {
	return currency.NullDecimalField{
		Valid:        true,
		DecimalField: currency.NewDecimalField(currency.FromE8Value(e8)),
	}, nil
}

func PairToProto(p currency.Pair) (string, error) {
	return p.String(), nil
}

func PairFromProto(p string) (currency.Pair, error) {
	return currency.StringToPair(p)
}
