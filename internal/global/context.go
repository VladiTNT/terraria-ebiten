package global

type Context struct {
	Width  int
	Height int
}

func NewContext() *Context {
	return &Context{
		Width:  640,
		Height: 360,
	}
}
