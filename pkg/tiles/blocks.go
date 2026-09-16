package tiles

import (
	"github.com/VladiTNT/terraria-ebiten/pkg/sprites"
	"github.com/hajimehoshi/ebiten/v2"
)

type Block uint8

const (
	Air Block = iota
	Dirt
)

var BlockAtlas = map[Block]*ebiten.Image{
	Dirt: sprites.Dirt,
}
