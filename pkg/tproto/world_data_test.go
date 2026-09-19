package tproto_test

import (
	"bytes"
	"testing"

	"github.com/VladiTNT/terraria-ebiten/pkg/tiles"
	"github.com/VladiTNT/terraria-ebiten/pkg/tproto"
)

func TestWorldData(t *testing.T) {
	conn := new(bytes.Buffer)
	var wd tproto.WorldData
	for i := range tproto.WorldLength {
		for j := range tproto.WorldHeight {
			wd.Blocks[i][j] = tiles.Dirt
		}
	}

	err := tproto.WritePacket(conn, tproto.NewPacket(tproto.Ping, tproto.WorldDataPayload(wd)))
	if err != nil {
		t.Errorf("Error writting packet to conn: %v\n", err)
	}

	p, err := tproto.ReadPacket(conn)
	if err != nil {
		t.Errorf("Error reading packet from conn: %V\n", err)
	}

	if p.Type != tproto.Ping {
		t.Errorf("Error wrong packet type: wanted %d, got %d\n", tproto.Ping, p.Type)
	}

	newWd := tproto.DecodeWorldDataPayload(p.Payload)

	if newWd.Blocks != wd.Blocks {
		t.Errorf("Error, world data doesn't match.\n")
	}
}
