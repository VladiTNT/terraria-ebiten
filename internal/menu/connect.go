package menu

import (
	"image/color"

	"github.com/VladiTNT/terraria-ebiten/internal/cnet"
	"github.com/VladiTNT/terraria-ebiten/internal/global"
	"github.com/VladiTNT/terraria-ebiten/pkg/txt"
	"github.com/VladiTNT/terraria-ebiten/pkg/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	UrlInputOffsetX = 70
)

type ConnectOption int

const (
	UrlInput ConnectOption = iota
	Join
)

type ConnectWindow struct {
	Context *global.Context

	UrlTextbox *ui.TextBox

	Options       []*ui.Label
	CurrentOption ConnectOption

	LoadingAnim LoadingAnimation
}

func NewConnectWindow(ctx *global.Context) *ConnectWindow {
	return &ConnectWindow{
		Context: ctx,

		UrlTextbox: ui.NewTextBox(
			MenuOptionX+WindowOffsetX+UrlInputOffsetX,
			1*MenuOptionVerticalSpacing, 40,
		),

		Options: []*ui.Label{
			ui.NewLabel(MenuOptionX+WindowOffsetX, 1*MenuOptionVerticalSpacing, "Url:"),
			ui.NewLabel(MenuOptionX+WindowOffsetX, 2*MenuOptionVerticalSpacing, "Join"),
		},

		LoadingAnim: NewLoadingAnimation(),
	}
}

func (cw *ConnectWindow) Update() {

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		if cw.CurrentOption > UrlInput {
			cw.CurrentOption -= 1
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		if cw.CurrentOption < Join {
			cw.CurrentOption += 1
		}
	}

	// URL text box active only when selected
	if cw.CurrentOption == UrlInput {
		cw.UrlTextbox.Update()
	}

	// Update the animation when we are loading
	if cw.Context.NetEngine.Status == cnet.Connecting {
		cw.LoadingAnim.Update()
	}

}

func (cw *ConnectWindow) Draw(screen *ebiten.Image, pp *txt.Printer) {
	// Draw URL text box
	if cw.CurrentOption == UrlInput {
		cw.UrlTextbox.Color = ui.ColorGREEN
		cw.UrlTextbox.Draw(screen, pp)
	} else {
		cw.UrlTextbox.Color = color.White
		cw.UrlTextbox.Draw(screen, pp)
	}

	// Options
	for i := range cw.Options {
		if i == int(cw.CurrentOption) {
			cw.Options[i].Draw(screen, pp, ui.ColorGREEN)
		} else {
			cw.Options[i].Draw(screen, pp, color.White)
		}
	}
}
