package server

import (
	"github.com/VladiTNT/terraria-ebiten/pkg/tiles"
	"github.com/VladiTNT/terraria-ebiten/pkg/tproto"
)

func GenerateWorld() tproto.WorldData {
	var wd tproto.WorldData

	for i := range tproto.WorldLength {
		for j := range tproto.WorldHeight {
			wd.Blocks[i][j] = tiles.Dirt
		}
	}

	return wd
}
