package interfaces

import "net"

type Handler interface {
	Handler()
	Disconnect()
	MakeHandler(conn net.Conn)
}
