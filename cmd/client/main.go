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
	// Window settings
	ebiten.SetWindowSize(DefaultWindowWidth, DefaultWindowHeight)
	ebiten.SetWindowTitle(WindowTitle)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
