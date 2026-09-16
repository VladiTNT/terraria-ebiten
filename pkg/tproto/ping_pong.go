package tproto

import (
	"bytes"
	"encoding/binary"
)

func PingPongPayload(x int64) []byte {
	return binary.AppendVarint([]byte{}, x)
}

func DecodePingPongPayload(buf []byte) (int64, error) {
	return binary.ReadVarint(bytes.NewReader(buf))
}
