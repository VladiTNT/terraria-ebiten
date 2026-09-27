package tproto_test

import (
	"testing"

	"github.com/VladiTNT/terraria-ebiten/pkg/tiles"
	"github.com/VladiTNT/terraria-ebiten/pkg/tproto"
)

func TestBlockChange(t *testing.T) {
	bc := tproto.BlockChange{2, 2, tiles.Dirt}

	newBc, err := tproto.DecodeBlockChangePayload(tproto.BlockChangePayload(bc))
	if err != nil {
		t.Errorf("Couldn't decode block change payload: %v\n", err)
	}

	if bc != newBc {
		t.Errorf("Block changes don't match: expected %v, got %v\n", bc, newBc)
	}
}
