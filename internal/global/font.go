package global

import (
	"github.com/VladiTNT/terraria-ebiten/assets"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func GetFont() *text.GoTextFaceSource {
	f, err := assets.FS.Open("fonts/AnnotationMNerdFont-Bold.ttf")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fnt, err := text.NewGoTextFaceSource(f)
	if err != nil {
		panic(err)
	}

	return fnt
}
