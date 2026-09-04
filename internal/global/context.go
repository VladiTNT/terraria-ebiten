package global

type Context struct {
	Width  int
	Height int

	Sprites *Sprites
}

func NewContext() *Context {
	return &Context{
		Width:  640,
		Height: 360,

		Sprites: NewSprites(),
	}
}
