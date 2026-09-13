package tproto

import (
	"encoding/binary"
	"errors"
	"io"
)

type PacketType uint16

const (
	Err PacketType = iota
	Ping
	Pong
)

type Packet struct {
	Type    PacketType
	Payload []byte
}

func WritePacket(w io.Writer, p Packet) error {
	err := binary.Write(w, binary.BigEndian, p.Type)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, uint32(len(p.Payload)))
	if err != nil {
		return err
	}

	n, err := w.Write(p.Payload)
	if err != nil {
		return err
	}

	if n != len(p.Payload) {
		return errors.New("tproto: mismatch between bytes written to conn and payload length")
	}

	return nil
}

func ReadPacket(r io.Reader) (Packet, error) {
	var id PacketType
	err := binary.Read(r, binary.BigEndian, &id)
	if err != nil {
		return Packet{}, err
	}

	var pLen uint32
	err = binary.Read(r, binary.BigEndian, &pLen)
	if err != nil {
		return Packet{}, err
	}

	buf := make([]byte, pLen)

	n, err := r.Read(buf)
	if err != nil {
		return Packet{}, err
	}

	if uint32(n) != pLen {
		return Packet{}, errors.New("tproto: mismatch between bytes read from conn and payload length")
	}

	return Packet{id, buf}, nil
}
