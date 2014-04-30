// handles conceptual areas within the game

package world

import (
	"github.com/streamweaver/server"
)

type Area struct {
	Entity
	occupants map[server.User]bool
	exits     []Exit


func (self *Area) Broadcast(line string) {
	for u, _ := range self.occupants {
		u.message <- line
	}
}
