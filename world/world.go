// Contains code for all in game entities and objects.
// file pogomud/unierse/universe.go

package world

type World struct {
	Areas   map[int]Area
	Players map[int]Player
}

func NewWorld() *World {
	w := new(World)
	w.Areas = make(map[int]Area)
	w.Players = make(map[int]Player)
	return w
}
