package protos

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

func ConvertMilliToTimestamp(millis int64) *timestamp.Timestamp {
	return &timestamp.Timestamp{
		Seconds: millis / 1000,
		Nanos:   int32((millis % 1000) * 1000000),
	}
}
