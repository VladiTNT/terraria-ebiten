package global

import (
	"github.com/VladiTNT/terraria-ebiten/internal/cnet"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Context struct {
	Width  int
	Height int

	NetEngine *cnet.Engine

	Sprites *Sprites
	Font    *text.GoTextFaceSource
}

func NewContext() *Context {
	return &Context{
		Width:  1280,
		Height: 720,

		NetEngine: cnet.NewEngine(),

		Sprites: NewSprites(),
		Font:    GetFont(),
	}
}
