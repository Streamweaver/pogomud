// Contains code for all in game entities and objects.
// file pogomud/unierse/universe.go

package world

// NOTE Stubs for stuff right now, will move to seperate files later.
type Entity struct {
	id          int
	name        string
	description string
}

type Room struct {
	Entity
}

type Character struct {
	Entity
}

type Item struct {
	Entity
}
