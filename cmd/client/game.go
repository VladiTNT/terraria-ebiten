package main

import (
	"github.com/VladiTNT/terraria-ebiten/internal/global"
	"github.com/VladiTNT/terraria-ebiten/internal/menu"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Context      *global.Context
	CurrentScene global.Scene
}

func NewGame() *Game {
	ctx := global.NewContext()

	return &Game{
		Context:      ctx,
		CurrentScene: menu.New(ctx),
	}
}

func (g *Game) Update() error {
	if scene := g.CurrentScene.Jump(); scene != nil {
		g.CurrentScene = scene
	}

	return g.CurrentScene.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.CurrentScene.Draw(screen)
}

func (g *Game) Layout(w, h int) (int, int) {
	return g.Context.Width, g.Context.Height
}
