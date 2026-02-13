package netcode

import (
	"encoding/gob"
	"net"
)

type NetClient struct {
	enc *gob.Encoder
	dec *gob.Decoder
}

func NewClient(addr string) (*NetClient, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &NetClient{
		enc: gob.NewEncoder(conn),
		dec: gob.NewDecoder(conn),
	}, nil
}

func (c *NetClient) SendInput(in Input) {
	c.enc.Encode(in)
}

func (c *NetClient) Receive(state *WorldState) error {
	return c.dec.Decode(state)
}
