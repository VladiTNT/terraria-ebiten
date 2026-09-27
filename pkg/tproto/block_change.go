package tproto

import (
	"encoding/binary"
	"fmt"

	"github.com/VladiTNT/terraria-ebiten/pkg/tiles"
)

type BlockChange struct {
	X, Y  uint64
	Block tiles.Block
}

func BlockChangePayload(bc BlockChange) []byte {
	var buff []byte

	buff = binary.BigEndian.AppendUint64(buff, bc.X)
	buff = binary.BigEndian.AppendUint64(buff, bc.Y)
	buff = append(buff, byte(bc.Block))

	return buff
}

func DecodeBlockChangePayload(buf []byte) (BlockChange, error) {
	if len(buf) != 17 {
		return BlockChange{}, fmt.Errorf("Block change payload is %d bytes instead of 17", len(buf))
	}

	return BlockChange{
		X:     binary.BigEndian.Uint64(buf[0:8]),
		Y:     binary.BigEndian.Uint64(buf[8:16]),
		Block: tiles.Block(buf[16]),
	}, nil
}
