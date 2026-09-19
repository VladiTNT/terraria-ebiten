package game

import (
	"github.com/VladiTNT/terraria-ebiten/internal/global"
	"github.com/VladiTNT/terraria-ebiten/pkg/netutils"
	"github.com/VladiTNT/terraria-ebiten/pkg/tiles"
	"github.com/VladiTNT/terraria-ebiten/pkg/tproto"
	"github.com/hajimehoshi/ebiten/v2"
)

type World struct {
	Alive   bool
	Context *global.Context

	WorldIsGood bool
	WorldMap    tproto.WorldData
}

func NewWorld(ctx *global.Context) *World {
	return &World{
		Alive:   true,
		Context: ctx,

		WorldIsGood: false,
		WorldMap:    tproto.WorldData{},
	}
}

func (w *World) Update() error {
	// Updates from the server
	for _, packet := range netutils.Drain(w.Context.NetEngine.ReadChan) {
		switch packet.Type {
		case tproto.WorldDataResponse:
			w.WorldMap = tproto.DecodeWorldDataPayload(packet.Payload)
			w.WorldIsGood = true
		}
	}

	if !w.WorldIsGood {
		w.Context.NetEngine.WriteChan <- tproto.NewPacket(tproto.WorldDataRequest, []byte{})
	}

	return nil
}

func (w *World) Draw(screen *ebiten.Image) {
	if w.WorldIsGood {
		// Loop through the world matrix
		for i := range tproto.WorldLength {
			for j := range tproto.WorldHeight {
				// Skip draw call for air tiles
				if w.WorldMap.Blocks[i][j] == tiles.Air {
					continue
				}

				op := new(ebiten.DrawImageOptions)
				op.GeoM.Translate(float64(i*tiles.BlockSideLength), float64(j*tiles.BlockSideLength))

				screen.DrawImage(tiles.BlockAtlas[w.WorldMap.Blocks[i][j]], op)
			}
		}
	}
}

func (w *World) Jump() global.Scene {
	if w.Alive {
		return nil
	}

	return nil
}
