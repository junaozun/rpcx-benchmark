package util

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
)

type GogoProtoBuf struct {
}

func (c *GogoProtoBuf) Decode(data []byte, i interface{}) error {
	if m, ok := i.(proto.Unmarshaler); ok {
		return m.Unmarshal(data)
	}
	if m, ok := i.(proto.Message); ok {
		return proto.Unmarshal(data, m)
	}
	return fmt.Errorf("%T is not a gogoproto.Unmarshaler  or gogopb.Message", i)
}

func (c *GogoProtoBuf) Encode(i interface{}) ([]byte, error) {
	if m, ok := i.(proto.Marshaler); ok {
		return m.Marshal()
	}

	if m, ok := i.(proto.Message); ok {
		return proto.Marshal(m)
	}

	return nil, fmt.Errorf("%T is not a gogoproto.Marshaler or gogopb.Message", i)
}
