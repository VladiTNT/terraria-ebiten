package cnet

import (
	"net"
	"time"

	"github.com/VladiTNT/terraria-ebiten/pkg/netutils"
)

const (
	// 60 Ticks per second
	DefaultTickRate = time.Second / 60
)

type Engine struct {
	ticker *time.Ticker
	conn   net.Conn

	errBuffer chan error
}

func NewEngine() *Engine {
	return &Engine{
		ticker: time.NewTicker(DefaultTickRate),
		conn:   nil,

		errBuffer: make(chan error, 10),
	}
}

func (e *Engine) Connect(addr string) {
	go func() {
		var err error

		e.conn, err = net.Dial("tcp", addr)
		if err != nil {
			e.errBuffer <- err
			// Return here to prevent crash if we get an error
			return
		}
	}()
}

func (e *Engine) Err() []error {
	return netutils.Drain(e.errBuffer)
}
