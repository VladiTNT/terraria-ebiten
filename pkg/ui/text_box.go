package ui

import (
	"image/color"

	"github.com/VladiTNT/terraria-ebiten/pkg/txt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type TextBox struct {
	X, Y  float64
	Text  string
	Color color.Color

	runeBuffer []rune
	maxLen     int
}

func NewTextBox(x, y float64, maxLen int) *TextBox {
	return &TextBox{
		X: x, Y: y,
		Text:  "",
		Color: color.White,

		runeBuffer: make([]rune, 10),
		maxLen:     maxLen,
	}
}

func (tb *TextBox) Update() {
	tb.runeBuffer = ebiten.AppendInputChars(tb.runeBuffer[:0])

	if len(tb.Text) < tb.maxLen {
		tb.Text += string(tb.runeBuffer)
	}

	if repeatingKeyPressed(ebiten.KeyBackspace) {
		if len(tb.Text) >= 1 {
			tb.Text = tb.Text[:len(tb.Text)-1]
		}
	}
}

func (tb *TextBox) Draw(screen *ebiten.Image, pp *txt.Printer) {
	pp.PrintWithPositionAndColor(screen, tb.Text+"_", tb.X, tb.Y, tb.Color)
}

// Stolen from ebiten website btw.
func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)

	d := inpututil.KeyPressDuration(key)

	if d == 1 {
		return true
	}

	if d >= delay && (d-delay)%interval == 0 {
		return true
	}

	return false
}
