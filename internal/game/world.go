package game

import (
	"github.com/VladiTNT/terraria-ebiten/internal/global"
	"github.com/VladiTNT/terraria-ebiten/pkg/tproto"
	"github.com/hajimehoshi/ebiten/v2"
)

type World struct {
	Alive   bool
	Context *global.Context

	WorldMap *tproto.WorldData
}

func NewWorld(ctx *global.Context) *World {
	return &World{
		Alive:   true,
		Context: ctx,

		WorldMap: nil,
	}
}

func (w *World) Update() error {

	return nil
}

func (w *World) Draw(screen *ebiten.Image) {

}

func (w *World) Jump() global.Scene {
	if w.Alive {
		return nil
	}

	return nil
}
