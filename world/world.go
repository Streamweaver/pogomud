// Contains code for all in game entities and objects.
// file pogomud/unierse/universe.go

package world

type World struct {
	Areas   map[int]Area
	Players map[int]Player
}

func NewWorld() *World {
	w := new(World)
	w.areas = make(map[int]Area)
	return w
}
