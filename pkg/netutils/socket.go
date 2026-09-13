package netutils

import (
	"bufio"
	"net"
	"time"

	"github.com/VladiTNT/terraria-ebiten/pkg/tproto"
)

const DefaultTickRate = time.Second / 60

func NewSockets(conn net.Conn, errChan chan<- error, alive *bool) (<-chan tproto.Packet, chan<- tproto.Packet) {
	readChan := make(chan tproto.Packet, 10)
	writeChan := make(chan tproto.Packet, 10)

	// Read gororuine
	go func() {
		rd := bufio.NewReader(conn)

		for *alive {
			p, err := tproto.ReadPacket(rd)
			if err != nil {
				errChan <- err
				*alive = false
			}

			readChan <- p
		}
	}()

	// Write goroutine
	go func() {
		wr := bufio.NewWriter(conn)
		ticker := time.NewTicker(DefaultTickRate)
		defer ticker.Stop()

		for *alive {
			<-ticker.C

			packets := Drain(writeChan)
			for _, p := range packets {
				err := tproto.WritePacket(wr, p)
				if err != nil {
					errChan <- err
					*alive = false
				}
			}

			err := wr.Flush()
			if err != nil {
				errChan <- err
				*alive = false
			}
		}
	}()

	return readChan, writeChan
}
