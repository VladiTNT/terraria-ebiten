package txt

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Printer struct {
	Font     *text.GoTextFaceSource
	FontSize float64
}

func NewPrinter(font *text.GoTextFaceSource, fontSize float64) *Printer {
	return &Printer{Font: font, FontSize: fontSize}
}

func (pp *Printer) Print(screen *ebiten.Image, s string, opts *text.DrawOptions) {
	PrintDefault(screen, s, pp.Font, pp.FontSize, opts)
}

func (pp *Printer) PrintWithPosition(screen *ebiten.Image, s string, x, y float64) {
	PrintWithPosition(screen, s, pp.Font, pp.FontSize, x, y)
}

func (pp *Printer) PrintWithPositionAndColor(screen *ebiten.Image, s string,
	x, y float64, col color.Color,
) {
	PrintWithPositionAndColor(screen, s, pp.Font, pp.FontSize, x, y, col)
}
