package server

import (
	"fmt"
	"time"

	"github.com/VladiTNT/terraria-ebiten/internal/server/config"
	"github.com/VladiTNT/terraria-ebiten/pkg/netutils"
	"github.com/VladiTNT/terraria-ebiten/pkg/tproto"
)

const (
	GameTickRate = time.Second / 60
)

type Game struct {
	IsRunning  bool
	PlayerChan chan Player
	Players    []Player

	Settings *config.Settings
	World    tproto.WorldData
}

func NewGame(settings *config.Settings) *Game {
	return &Game{
		IsRunning:  true,
		PlayerChan: make(chan Player, 10),
		Players:    make([]Player, 0, 10),

		Settings: settings,
		World:    GenerateWorld(),
	}
}

func (g *Game) Main() {
	// 60 TPS Ticker
	ticker := time.NewTicker(GameTickRate)
	defer ticker.Stop()

	for range ticker.C {
		// If the game has to close we stop the loop.
		if !g.IsRunning {
			break
		}

		// Add new players
		g.Players = append(g.Players, netutils.Drain(g.PlayerChan)...)

		// Handle player requests
		for _, player := range g.Players {
			g.HandlePlayer(player)
		}
	}
}

func (g *Game) HandlePlayer(p Player) {
	// For each packet sent by the player
	for _, packet := range netutils.Drain(p.ReadChan) {
		// Switch to change action depending on packet type
		switch packet.Type {
		// Ping-Pong handling
		case tproto.Ping:
			n, err := tproto.DecodePingPongPayload(packet.Payload)
			if err != nil {
				fmt.Printf("Error decoding ping-pong payload: %v\n", err)
				continue
			}

			p.WriteChan <- tproto.NewPacket(tproto.Pong, tproto.PingPongPayload(n))
		// WorldData request handling
		case tproto.WorldDataRequest:
			p.WriteChan <- tproto.NewPacket(tproto.WorldDataResponse, tproto.WorldDataPayload(g.World))
		}
	}
}
