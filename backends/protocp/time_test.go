package protocp_test

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/luno/jettison/errors"
	"github.com/luno/jettison/jtest"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"bitx/backends/protocp"
)

func TestTimeToProto(t *testing.T) {
	secs := randRange(-62135596800, +253402300799)
	nanos := randRange(0, 1e9)

	tests := []struct {
		name   string
		t      time.Time
		pb     *timestamppb.Timestamp
		expErr error
	}{
		{
			name: "empty",
			t:    time.Time{},
			pb:   &timestamppb.Timestamp{Seconds: -62135596800},
		},
		{
			name: "zero",
			t:    time.Unix(0, 0).UTC(),
			pb:   &timestamppb.Timestamp{},
		},
		{
			name: "past",
			t:    time.Date(1989, time.November, 9, 8, 7, 6, 5, time.UTC),
			pb:   &timestamppb.Timestamp{Seconds: 626602026, Nanos: 5},
		},
		{
			name: "future",
			t:    time.Date(2089, time.November, 9, 8, 7, 6, 5, time.UTC),
			pb:   &timestamppb.Timestamp{Seconds: 3782362026, Nanos: 5},
		},
		{
			name: "random",
			t:    time.Unix(secs, nanos).UTC(),
			pb:   &timestamppb.Timestamp{Seconds: secs, Nanos: int32(nanos)},
		},
		{
			name:   "overflow",
			t:      time.Date(10089, time.November, 9, 8, 7, 6, 5, time.UTC),
			expErr: errors.New("invalid time.Time"),
		},
		{
			name:   "underflow",
			t:      time.Date(0, time.November, 9, 8, 7, 6, 5, time.UTC),
			expErr: errors.New("invalid time.Time"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cpb, err := protocp.TimeToProto(tc.t)
			jtest.Require(t, tc.expErr, err)
			require.Equal(t, tc.pb, cpb)
		})
	}
}

func TestTimeFromProto(t *testing.T) {
	secs := randRange(-62135596800, +253402300799)
	nanos := randRange(0, 1e9)

	tests := []struct {
		name   string
		pb     *timestamppb.Timestamp
		t      time.Time
		expErr error
	}{
		{
			name: "empty",
			pb:   &timestamppb.Timestamp{Seconds: -62135596800},
			t:    time.Time{},
		},
		{
			name: "zero",
			pb:   &timestamppb.Timestamp{},
			t:    time.Unix(0, 0).UTC(),
		},
		{
			name: "past",
			pb:   &timestamppb.Timestamp{Seconds: 626602026, Nanos: 5},
			t:    time.Date(1989, time.November, 9, 8, 7, 6, 5, time.UTC),
		},
		{
			name: "future",
			pb:   &timestamppb.Timestamp{Seconds: 3782362026, Nanos: 5},
			t:    time.Date(2089, time.November, 9, 8, 7, 6, 5, time.UTC),
		},
		{
			name: "random",
			pb:   &timestamppb.Timestamp{Seconds: secs, Nanos: int32(nanos)},
			t:    time.Unix(secs, nanos).UTC(),
		},
		{
			name:   "overflow secs",
			pb:     &timestamppb.Timestamp{Seconds: 999999999999, Nanos: 0},
			expErr: errors.New("invalid timestamppb.Timestamp"),
		},
		{
			name:   "overflow nanos",
			pb:     &timestamppb.Timestamp{Seconds: 12345, Nanos: 2e9},
			expErr: errors.New("invalid timestamppb.Timestamp"),
		},
		{
			name:   "underflow secs",
			pb:     &timestamppb.Timestamp{Seconds: -999999999999, Nanos: 0},
			expErr: errors.New("invalid timestamppb.Timestamp"),
		},
		{
			name:   "underflow nanos",
			pb:     &timestamppb.Timestamp{Seconds: 12345, Nanos: -5},
			expErr: errors.New("invalid timestamppb.Timestamp"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ct, err := protocp.TimeFromProto(tc.pb)
			jtest.Require(t, tc.expErr, err)
			require.Equal(t, tc.t, ct)
		})
	}
}

func TestDurationToProto(t *testing.T) {
	nanos := rand.Int63()
	secs := nanos / 1e9
	nanos = nanos % 1e9

	tests := []struct {
		name string
		d    time.Duration
		pb   *durationpb.Duration
	}{
		{
			name: "zero",
			d:    time.Duration(0),
			pb:   &durationpb.Duration{},
		},
		{
			name: "neg",
			d:    -1234*time.Second + -5678*time.Nanosecond,
			pb:   &durationpb.Duration{Seconds: -1234, Nanos: -5678},
		},
		{
			name: "pos",
			d:    1234*time.Second + 5678*time.Nanosecond,
			pb:   &durationpb.Duration{Seconds: 1234, Nanos: 5678},
		},
		{
			name: "random",
			d:    time.Duration(secs)*time.Second + time.Duration(nanos)*time.Nanosecond,
			pb:   &durationpb.Duration{Seconds: secs, Nanos: int32(nanos)},
		},
		{
			name: "max",
			d:    time.Duration(math.MaxInt64),
			pb:   &durationpb.Duration{Seconds: 9223372036, Nanos: int32(854775807)},
		},
		{
			name: "min",
			d:    time.Duration(math.MinInt64),
			pb:   &durationpb.Duration{Seconds: -9223372036, Nanos: int32(-854775808)},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cpb, err := protocp.DurationToProto(tc.d)
			jtest.RequireNil(t, err)
			require.Equal(t, tc.pb, cpb)
		})
	}
}

func TestDurationFromProto(t *testing.T) {
	nanos := rand.Int63()
	secs := nanos / 1e9
	nanos = nanos % 1e9

	tests := []struct {
		name   string
		pb     *durationpb.Duration
		d      time.Duration
		expErr error
	}{
		{
			name: "zero",
			pb:   &durationpb.Duration{},
			d:    time.Duration(0),
		},
		{
			name: "neg",
			pb:   &durationpb.Duration{Seconds: -1234, Nanos: -5678},
			d:    -1234*time.Second + -5678*time.Nanosecond,
		},
		{
			name: "pos",
			pb:   &durationpb.Duration{Seconds: 1234, Nanos: 5678},
			d:    1234*time.Second + 5678*time.Nanosecond,
		},
		{
			name: "random",
			pb:   &durationpb.Duration{Seconds: secs, Nanos: int32(nanos)},
			d:    time.Duration(secs)*time.Second + time.Duration(nanos)*time.Nanosecond,
		},
		{
			name: "max",
			pb:   &durationpb.Duration{Seconds: 9223372036, Nanos: int32(854775807)},
			d:    time.Duration(math.MaxInt64),
		},
		{
			name: "min",
			pb:   &durationpb.Duration{Seconds: -9223372036, Nanos: int32(-854775808)},
			d:    time.Duration(math.MinInt64),
		},
		{
			name:   "overflow",
			pb:     &durationpb.Duration{Seconds: 9223372036000, Nanos: int32(854775807)},
			expErr: errors.New("invalid durationpb.Duration"),
		},
		{
			name:   "underflow",
			pb:     &durationpb.Duration{Seconds: -9223372036000, Nanos: int32(-854775808)},
			expErr: errors.New("invalid durationpb.Duration"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cd, err := protocp.DurationFromProto(tc.pb)
			jtest.Require(t, tc.expErr, err)
			require.Equal(t, tc.d, cd)
		})
	}
}

func randRange(min, max int64) int64 {
	return rand.Int63n(max-min) + min
}
