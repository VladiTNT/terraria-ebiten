package tproto_test

import (
	"bytes"
	"testing"

	"github.com/VladiTNT/terraria-ebiten/pkg/tproto"
)

func TestPacket(t *testing.T) {
	conn := new(bytes.Buffer)
	ids := []tproto.PacketType{tproto.Err, tproto.Ping, tproto.Pong}
	payloads := []string{"Nerds", "Birds", "C++ programming language"}

	for i := range 3 {
		err := tproto.WritePacket(conn, tproto.Packet{ids[i], []byte(payloads[i])})
		if err != nil {
			t.Errorf("Error writting packets to conn: %v\n", err)
		}
	}

	for i := range 3 {
		p, err := tproto.ReadPacket(conn)
		if err != nil {
			t.Errorf("Error reading packets to conn: %v\n", err)
		}

		if p.Type != ids[i] || !bytes.Equal(p.Payload, []byte(payloads[i])) {
			t.Errorf("Mismatch, wanted %d %s, got %d %s.\n", ids[i], payloads[i], p.Type, p.Payload)
		}
	}
}
