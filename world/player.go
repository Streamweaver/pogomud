package world

import (
	"net"
)

type Player struct {
	Entity
	password string
	conn     *net.TCPConn
	receive  *chan string
}
