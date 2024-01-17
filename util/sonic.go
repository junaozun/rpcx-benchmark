package util

import "github.com/bytedance/sonic"

type SonicJson struct {
}

func (c *SonicJson) Decode(data []byte, i interface{}) error {
	return sonic.Unmarshal(data, i)
}

func (c *SonicJson) Encode(i interface{}) ([]byte, error) {
	return sonic.Marshal(i)
}
