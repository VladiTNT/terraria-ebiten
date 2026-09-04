package menu

import (
	"github.com/VladiTNT/terraria-ebiten/internal/global"
	"github.com/hajimehoshi/ebiten/v2"
)

type Menu struct {
	Alive   bool
	Context *global.Context

	Message string
}

func New(ctx *global.Context) *Menu {
	return &Menu{
		Alive:   true,
		Context: ctx,

		Message: "Hello, world!",
	}
}

func (m *Menu) Update() error {
	return nil
}

func (m *Menu) Draw(screen *ebiten.Image) {
	for x := 0; x < 32; x += 8 {
		for y := 0; y < 32; y += 8 {
			op := new(ebiten.DrawImageOptions)
			op.GeoM.Translate(float64(x), float64(y))
			screen.DrawImage(m.Context.Sprites.Dirt, op)
		}
	}
}

func (m *Menu) Jump() global.Scene {
	return nil
}
