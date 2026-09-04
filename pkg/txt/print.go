package txt

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/text/language"
)

const (
	DefaultLineSpacingMultiplier float64 = 1.2
)

var (
	DefaultDrawingDirection = text.DirectionLeftToRight
	DefaultLanguage         = language.English
)

func Print(screen *ebiten.Image, s string,
	font *text.GoTextFaceSource, direction text.Direction, fontSize float64, lang language.Tag,
	opts *text.DrawOptions,
) {
	text.Draw(screen, s, &text.GoTextFace{
		Source:    font,
		Direction: direction,
		Size:      fontSize,
		Language:  lang,
	}, opts)
}

func PrintDefault(screen *ebiten.Image, s string,
	font *text.GoTextFaceSource, fontSize float64, opts *text.DrawOptions,
) {
	Print(screen, s, font, DefaultDrawingDirection, fontSize, DefaultLanguage, opts)
}

func PrintWithPosition(screen *ebiten.Image, s string,
	font *text.GoTextFaceSource, fontSize, x, y float64,
) {
	op := new(text.DrawOptions)
	op.LineSpacing = DefaultLineSpacingMultiplier * fontSize
	op.GeoM.Translate(x, y)

	PrintDefault(screen, s, font, fontSize, op)
}

func PrintWithPositionAndColor(screen *ebiten.Image, s string,
	font *text.GoTextFaceSource, fontSize, x, y float64, col color.Color,
) {
	op := new(text.DrawOptions)
	op.LineSpacing = DefaultLineSpacingMultiplier * fontSize
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(col)

	PrintDefault(screen, s, font, fontSize, op)
}
