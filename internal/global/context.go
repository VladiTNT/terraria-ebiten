package global

import "github.com/hajimehoshi/ebiten/v2/text/v2"

type Context struct {
	Width  int
	Height int

	Sprites *Sprites
	Font    *text.GoTextFaceSource
}

func NewContext() *Context {
	return &Context{
		Width:  1280,
		Height: 720,

		Sprites: NewSprites(),
		Font:    GetFont(),
	}
}
