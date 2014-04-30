// Handles things related to in world player characters.
package world

// NOTE Stubs for stuff right now, will move to seperate files later.
type Entity struct {
	id          int
	Name        string
	Description string
}

type Mob struct {
	Entity
	Strength     int
	Intelligence int
	Dexterity    int

	area *world.Area
}
