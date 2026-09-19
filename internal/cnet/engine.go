package cnet

import (
	"fmt"
	"net"

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

		e.Alive = true

		e.ReadChan, e.WriteChan = netutils.NewSockets(e.conn, &e.Alive)

		e.Status = Connected
	}()
}
