package server

import (
	"fmt"
	"net"
)

func (s *Server) handleConn(conn net.Conn) {
	defer conn.Close()

	fmt.Fprintf(conn, "Hello noob")

	select {}
}
