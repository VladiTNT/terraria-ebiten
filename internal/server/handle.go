package server

import (
	"net"

	"github.com/VladiTNT/terraria-ebiten/pkg/netutils"
)

func (s *Server) handleConn(conn net.Conn) {
	defer conn.Close()

	alive := true
	p := NewPlayer(netutils.NewSockets(conn, &alive))

	s.Game.PlayerChan <- p

	select {}
}
