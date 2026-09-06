package menu

import (
	"fmt"

	"github.com/VladiTNT/terraria-ebiten/internal/global"
	"github.com/VladiTNT/terraria-ebiten/pkg/txt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type LoadingScreen struct {
	Alive   bool
	Context *global.Context

	FrameCounter int
	Frames       []string

	Text string
}

func NewLoadingScreen(ctx *global.Context) *LoadingScreen {
	return &LoadingScreen{
		Alive:   true,
		Context: ctx,

		FrameCounter: 0,
		Frames: []string{
			"Waiting",
			"Waiting.",
			"Waiting..",
			"Waiting...",
		},
	}
}

func (ls *LoadingScreen) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		ls.Alive = false
	}

	ls.FrameCounter++

	if ls.FrameCounter >= 120 {
		ls.FrameCounter = 0
	}

	// Check errors
	nErrs := ls.Context.NetEngine.Err()
	if len(nErrs) > 0 {
		for _, nErr := range nErrs {
			fmt.Println(nErr)
		}
	}

	s := ls.Context.NetEngine.Thing()
	if s != "" {
		ls.Text = s
	}

	return nil
}

func (ls *LoadingScreen) Draw(screen *ebiten.Image) {
	txt.PrintWithPosition(screen, ls.Frames[ls.FrameCounter/30], ls.Context.Font, 24, 0, 0)
	txt.PrintWithPosition(screen, ls.Text, ls.Context.Font, 24, 0, 24)
}

func (ls *LoadingScreen) Jump() global.Scene {
	if ls.Alive {
		return nil
	}

	return New(ls.Context)
}
