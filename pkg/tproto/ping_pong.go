package tproto

import (
	"encoding/binary"
)

func PingPongPayload(x int64) []byte {
	return binary.AppendVarint([]byte{}, x)
}

func DecodePingPongPayload(buf []byte) (int64, error) {
	var n int64
	_, err := binary.Decode(buf, binary.BigEndian, &n)
	if err != nil {
		return 0, err
	}
	return n, nil
}
