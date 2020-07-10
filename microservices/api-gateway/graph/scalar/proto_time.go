package scalar

import (
	"errors"
	"io"
	"strconv"
	"time"

	"github.com/ezio1119/fishapp-api-gateway/conf"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalTimeProto(protoT timestamp.Timestamp) graphql.Marshaler {
	t, _ := ptypes.Timestamp(&protoT)
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(t.In(conf.C.Sv.ClientLocation).Format(time.RFC3339)))
	})
}

func UnmarshalTimeProto(v interface{}) (timestamp.Timestamp, error) {
	if tmpStr, ok := v.(string); ok {
		t, err := time.Parse(time.RFC3339, tmpStr)
		if err != nil {
			return timestamp.Timestamp{}, err
		}

		a, err := ptypes.TimestampProto(t)
		if err != nil {
			return timestamp.Timestamp{}, err
		}

		return *a, nil
	}
	return timestamp.Timestamp{}, errors.New("time should be RFC3339 formatted string")
}
