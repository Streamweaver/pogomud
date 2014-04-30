// Contains code for all in game entities and objects.
// file pogomud/unierse/universe.go

package world

type World struct {
	areas map[Area]bool
}

func NewWorld() *World {
	w := new(World)
	w.areas = make(map[Area]bool)
	return w
}

func (w *World) AddArea(a *Area) {
	w.areas[a] = true
}

func (w *World) RemoveArea(a *Area) {
	delete(a, w.areas)
}
