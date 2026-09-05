package menu

import (
	"image/color"

	"github.com/VladiTNT/terraria-ebiten/internal/global"
	"github.com/VladiTNT/terraria-ebiten/pkg/txt"
	"github.com/VladiTNT/terraria-ebiten/pkg/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	MenuOptionX               = 40
	MenuOptionVerticalSpacing = 32
)

type MenuOption int

const (
	Connect MenuOption = iota
	Character
	Settings
	Quit
)

type Menu struct {
	Alive   bool
	Context *global.Context

	TextEngine *txt.Printer

	Options       []*ui.Label
	CurrentOption MenuOption
}

func New(ctx *global.Context) *Menu {
	return &Menu{
		Alive:   true,
		Context: ctx,

		TextEngine: txt.NewPrinter(ctx.Font, 24),

		Options: []*ui.Label{
			ui.NewLabel(MenuOptionX, 1*MenuOptionVerticalSpacing, "Connect"),
			ui.NewLabel(MenuOptionX, 2*MenuOptionVerticalSpacing, "Character"),
			ui.NewLabel(MenuOptionX, 3*MenuOptionVerticalSpacing, "Settings"),
			ui.NewLabel(MenuOptionX, 4*MenuOptionVerticalSpacing, "Quit"),
		},
		CurrentOption: Connect,
	}
}

func (m *Menu) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	// Moving up in the option list.
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		if m.CurrentOption > Connect {
			m.CurrentOption -= 1
		}
	}

	// Moving down in the option list
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		if m.CurrentOption < Quit {
			m.CurrentOption += 1
		}
	}

	return nil
}

func (m *Menu) Draw(screen *ebiten.Image) {
	// Draw Menu options
	for i := range m.Options {
		if i == int(m.CurrentOption) {
			m.Options[i].Draw(screen, m.TextEngine, ui.ColorRED)
		} else {
			m.Options[i].Draw(screen, m.TextEngine, color.White)
		}
	}
}

func (m *Menu) Jump() global.Scene {
	return nil
}
