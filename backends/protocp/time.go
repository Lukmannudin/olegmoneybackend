package protocp

import (
	"time"

	"github.com/luno/jettison/errors"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TimeToProto(t time.Time) (*timestamppb.Timestamp, error) {
	pb := timestamppb.New(t)
	err := pb.CheckValid()
	if err != nil {
		return nil, errors.Wrap(err, "invalid time.Time")
	}
	return pb, nil
}

func TimeFromProto(t *timestamppb.Timestamp) (time.Time, error) {
	err := t.CheckValid()
	if err != nil {
		return time.Time{}, errors.Wrap(err, "invalid timestamppb.Timestamp")
	}
	return t.AsTime(), nil
}

func DurationToProto(d time.Duration) (*durationpb.Duration, error) {
	// No need to .CheckValid() because its impossible to construct an invalid
	// durationpb.Duration from a time.Duration.
	return durationpb.New(d), nil
}

func DurationFromProto(d *durationpb.Duration) (time.Duration, error) {
	err := d.CheckValid()
	if err != nil {
		return 0, errors.Wrap(err, "invalid durationpb.Duration")
	}
	return d.AsDuration(), nil
}

func TimeToProtoMs(t time.Time) (int64, error) {
	return t.UnixNano() / 1e6, nil
}

func TimeFromProtoMs(ms int64) (time.Time, error) {
	return time.Unix(0, ms*1e6), nil
}
