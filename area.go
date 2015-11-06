// handles conceptual areas within the game

package world

type Area struct {
	Entity
	occupants map[Player]bool
}

func (self *Area) Broadcast(line string) {
	
}
