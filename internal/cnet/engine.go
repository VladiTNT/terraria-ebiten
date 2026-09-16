package cnet

import (
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/VladiTNT/terraria-ebiten/pkg/netutils"
	"github.com/VladiTNT/terraria-ebiten/pkg/tproto"
)

type NetStatus int

const (
	NoConnection NetStatus = iota
	Connecting
	Connected
)

type Engine struct {
	Alive     bool
	Status    NetStatus
	ReadChan  <-chan tproto.Packet
	WriteChan chan<- tproto.Packet

	conn net.Conn
}

func NewEngine() *Engine {
	return &Engine{
		Alive:     false,
		Status:    NoConnection,
		ReadChan:  nil,
		WriteChan: nil,

		conn: nil,
	}
}

func (e *Engine) Connect(addr string) {
	go func() {
		e.Status = Connecting

		var err error

		e.conn, err = net.Dial("tcp", addr)
		if err != nil {
			fmt.Printf("Error dialing %s: %v\n", addr, err)
			return
		}

		n := rand.Int63()
		start := time.Now()
		err = tproto.WritePacket(e.conn, tproto.NewPacket(tproto.Ping, tproto.PingPongPayload(n)))
		if err != nil {
			fmt.Printf("Error writting packet: %v\n", err)
			e.conn.Close()
			e.Status = NoConnection
			return
		}

		p, err := tproto.ReadPacket(e.conn)
		if err != nil {
			fmt.Printf("Error reading packet: %v\n", err)
			e.conn.Close()
			e.Status = NoConnection
			return
		}

		if p.Type != tproto.Pong {
			fmt.Println("Error, didn't receive pong from server.")
		}

		newN, err := tproto.DecodePingPongPayload(p.Payload)
		if err != nil {
			fmt.Printf("Error decoding ping-pong: %v\n", err)
		}

		if n != newN {
			fmt.Printf("Error, number mutated during roundtrip: %d -> %d.\n", n, newN)
		}

		fmt.Printf("Ping: %v\n", time.Since(start))

		e.Alive = true

		e.ReadChan, e.WriteChan = netutils.NewSockets(e.conn, &e.Alive)

		e.Status = Connected
	}()
}
