package server

import (
	"fmt"
	"net"
	"time"

	"github.com/VladiTNT/terraria-ebiten/pkg/netutils"
	"github.com/VladiTNT/terraria-ebiten/pkg/tproto"
)

func (s *Server) handleConn(conn net.Conn) {
	defer conn.Close()

	alive := true
	ticker := time.NewTicker(s.Config.GameTickInterval)
	defer ticker.Stop()

	readChan, writeChan := netutils.NewSockets(conn, &alive)

	for range ticker.C {
		if !alive {
			break
		}

		for _, p := range netutils.Drain(readChan) {
			switch p.Type {
			case tproto.Ping:
				n, err := tproto.DecodePingPongPayload(p.Payload)
				if err != nil {
					fmt.Printf("Error decoding ping-pong: %v\n", err)
				}

				fmt.Println(n)

				writeChan <- tproto.NewPacket(tproto.Pong, tproto.PingPongPayload(n))
			}
		}
	}
}
