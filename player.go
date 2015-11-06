package world

import (
	"net"
	"fmt"
)

type Player struct {
	Entity
	password string
	conn     *net.TCPConn
	receive  *chan string
}

func FindPlayer(name string) (*Player, error) {
	return nil, fmt.Errorf("'%s' not found.", name)
}
