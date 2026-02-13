package netcode

import (
	"encoding/gob"
	"net"
	"sync"
)

type Client struct {
	conn  net.Conn
	enc   *gob.Encoder
	dec   *gob.Decoder
	Input Input
}

type Server struct {
	Clients []*Client
	mu      sync.Mutex
}

func NewServer(port string) (*Server, error) {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return nil, err
	}

	s := &Server{}

	go func() {
		for {
			conn, _ := ln.Accept()

			c := &Client{
				conn: conn,
				enc:  gob.NewEncoder(conn),
				dec:  gob.NewDecoder(conn),
			}

			s.mu.Lock()
			s.Clients = append(s.Clients, c)
			s.mu.Unlock()

			go s.readLoop(c)
		}
	}()
	return s, nil
}

func (s *Server) readLoop(c *Client) {
	for {
		var in Input
		if err := c.dec.Decode(&in); err != nil {
			return
		}
		c.Input = in
	}
}

func (s *Server) Broadcast(state WorldState) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, c := range s.Clients {
		c.enc.Encode(state)
	}
}
