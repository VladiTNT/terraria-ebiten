package sprites

import (
	"image/png"

	"github.com/VladiTNT/terraria-ebiten/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

func NewSprite(path string) *ebiten.Image {
	f, err := assets.FS.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
