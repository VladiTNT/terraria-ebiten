package tproto

import (
	"github.com/VladiTNT/terraria-ebiten/pkg/tiles"
)

const (
	WorldLength = 10
	WorldHeight = 5
)

type WorldData struct {
	Blocks [WorldLength][WorldHeight]tiles.Block
}

func WorldDataPayload(wd WorldData) []byte {
	var buf = make([]byte, 0, 8000000)

	for i := range WorldLength {
		for j := range WorldHeight {
			buf = append(buf, byte(wd.Blocks[i][j]))
		}
	}

	return buf
}

func DecodeWorldDataPayload(buf []byte) WorldData {
	var wd WorldData

	for i := range WorldLength {
		for j := range WorldHeight {
			wd.Blocks[i][j] = tiles.Block(buf[i+j])
		}
	}

	return wd
}
