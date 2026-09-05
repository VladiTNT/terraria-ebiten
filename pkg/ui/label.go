package ui

import (
	"image/color"

	"github.com/VladiTNT/terraria-ebiten/pkg/txt"
	"github.com/hajimehoshi/ebiten/v2"
)

type Label struct {
	X, Y float64
	Text string
}

func NewLabel(x, y float64, s string) *Label {
	return &Label{
		X: x, Y: y,
		Text: s,
	}
}

func (l *Label) Draw(screen *ebiten.Image, pp *txt.Printer, col color.Color) {
	pp.PrintWithPositionAndColor(screen, l.Text, l.X, l.Y, col)
}
