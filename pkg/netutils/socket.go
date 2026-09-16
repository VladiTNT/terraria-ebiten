package netutils

import (
	"bufio"
	"fmt"
	"net"

	"github.com/VladiTNT/terraria-ebiten/pkg/tproto"
)

func NewSockets(conn net.Conn, alive *bool) (<-chan tproto.Packet, chan<- tproto.Packet) {
	readChan := make(chan tproto.Packet, 10)
	writeChan := make(chan tproto.Packet, 10)

	// Read gororuine
	go func() {
		rd := bufio.NewReader(conn)

		for *alive {
			p, err := tproto.ReadPacket(rd)
			if err != nil {
				fmt.Printf("Error reading packet: %v\n", err)
				*alive = false
			}

			readChan <- p
		}
	}()

	// Write goroutine
	go func() {
		for *alive {
			err := tproto.WritePacket(conn, <-writeChan)
			if err != nil {
				fmt.Printf("Error writting packet: %v\n", err)
				*alive = false
			}
		}
	}()

	return readChan, writeChan
}
