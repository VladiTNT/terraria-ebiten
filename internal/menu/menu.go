package menu

import (
	"github.com/VladiTNT/terraria-ebiten/internal/global"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
	ebitenutil.DebugPrint(screen, m.Message)
}

func (m *Menu) Jump() global.Scene {
	return nil
}
