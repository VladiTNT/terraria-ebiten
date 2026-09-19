package server

import (
	"github.com/VladiTNT/terraria-ebiten/pkg/tproto"
)

type Player struct {
	ReadChan  <-chan tproto.Packet
	WriteChan chan<- tproto.Packet
}

func NewPlayer(rc <-chan tproto.Packet, wc chan<- tproto.Packet) Player {
	return Player{
		ReadChan:  rc,
		WriteChan: wc,
	}
}
