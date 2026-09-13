package global

import (
	"errors"
	"fmt"
	"net"

	"github.com/VladiTNT/terraria-ebiten/internal/cnet"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Context struct {
	Width  int
	Height int

	ErrChan   chan error
	NetEngine *cnet.Engine

	Sprites *Sprites
	Font    *text.GoTextFaceSource
}

func NewContext() *Context {
	ctx := new(Context)

	ctx.Width = 1280
	ctx.Height = 720

	ctx.ErrChan = make(chan error, 10)
	ctx.NetEngine = cnet.NewEngine(ctx.ErrChan)

	ctx.Sprites = NewSprites()
	ctx.Font = GetFont()

	// Global background jobs
	ctx.errHandler()

	return ctx
}

func (ctx *Context) errHandler() {
	go func() {
		for {
			err := <-ctx.ErrChan

			var netErr net.Error
			if errors.As(err, &netErr) {
				if ctx.NetEngine.Status != cnet.NoConnection {
					ctx.NetEngine.Alive = false
					ctx.NetEngine.Status = cnet.NoConnection
				}
			}

			fmt.Printf("Error: %v\n", err)
		}
	}()
}
