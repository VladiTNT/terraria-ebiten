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

	WindowOffsetX = 200
)

type MenuOption int

const (
	Connect MenuOption = iota
	Characters
	Settings
	Quit
)

type WindowSelect int

const (
	NoWindow WindowSelect = iota
	ConnWindow
	SettWindow
)

type Menu struct {
	Alive   bool
	Context *global.Context

	TextEngine *txt.Printer

	Options       []*ui.Label
	CurrentOption MenuOption

	CurrentWindow WindowSelect
	ConnectWindow *ConnectWindow
}

func New(ctx *global.Context) *Menu {
	return &Menu{
		Alive:   true,
		Context: ctx,

		TextEngine: txt.NewPrinter(ctx.Font, 24),

		Options: []*ui.Label{
			ui.NewLabel(MenuOptionX, 1*MenuOptionVerticalSpacing, "Connect"),
			ui.NewLabel(MenuOptionX, 2*MenuOptionVerticalSpacing, "Settings"),
			ui.NewLabel(MenuOptionX, 3*MenuOptionVerticalSpacing, "Characters"),
			ui.NewLabel(MenuOptionX, 4*MenuOptionVerticalSpacing, "Quit"),
		},
		CurrentOption: Connect,

		CurrentWindow: NoWindow,
		ConnectWindow: NewConnectWindow(),
	}
}

func (m *Menu) Update() error {
	// Close game only when no side window
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		if m.CurrentWindow == NoWindow {
			return ebiten.Termination
		} else {
			m.CurrentWindow = NoWindow
		}
	}

	// Handling input inside of sub windows
	if m.CurrentWindow != NoWindow {
		switch m.CurrentWindow {
		case ConnWindow:
			m.ConnectWindow.Update()
		}
	} else {
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
	}

	// Pressing enter logic
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		switch m.CurrentWindow {
		case NoWindow:
			switch m.CurrentOption {
			case Connect:
				m.CurrentWindow = ConnWindow
			case Quit:
				return ebiten.Termination
			}
		case ConnWindow:
			switch m.ConnectWindow.CurrentOption {
			// When joining a server
			case Join:
				// Connect to server
				m.Context.NetEngine.Connect(m.ConnectWindow.UrlTextbox.Text)
				// Close main menu
				m.Alive = false
			}
		}
	}

	return nil
}

func (m *Menu) Draw(screen *ebiten.Image) {
	// Draw Menu options
	for i := range m.Options {
		if i == int(m.CurrentOption) {
			m.Options[i].Draw(screen, m.TextEngine, ui.ColorGREEN)
		} else {
			m.Options[i].Draw(screen, m.TextEngine, color.White)
		}
	}

	// Draw window
	switch m.CurrentWindow {
	case ConnWindow:
		m.ConnectWindow.Draw(screen, m.TextEngine)
	}
}

func (m *Menu) Jump() global.Scene {
	if m.Alive {
		return nil
	}

	return NewLoadingScreen(m.Context)
}
