package global

import (
	"github.com/VladiTNT/terraria-ebiten/internal/cnet"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Context struct {
	Width  int
	Height int

	NetEngine *cnet.Engine

	Font *text.GoTextFaceSource
}

func NewContext() *Context {
	ctx := new(Context)

	ctx.Width = 1280
	ctx.Height = 720

	ctx.NetEngine = cnet.NewEngine()

	ctx.Font = GetFont()

	return ctx
}
