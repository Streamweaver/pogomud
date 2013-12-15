// handles conceptual areas within the game

package world

type Area struct {
	Entity
	occupants CharacterRegistry
	exits     map[string]Exit
}

func (self *Area) Enter(c Character) {
	self.Broadcast(c.name + " enters.")
	self.occupants.Add(c)
}

func (self *Area) Leave(c) {
	self.Broadcast(c.name + " leaves.")
	self.occupants.Remove(c)
}

func (self *Area) Broadcast(line string) {
	for _, c := range self.occupants.iMap {
		c.message <- line
	}
}

type Exit struct {
	Entity
	destID int //
}
