package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	DefaultWindowWidth  = 1280
	DefaultWindowHeight = 720
	WindowTitle         = "Terraria"
)

func main() {
	ebiten.SetWindowSize(DefaultWindowWidth, DefaultWindowHeight)
	ebiten.SetWindowTitle(WindowTitle)
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
