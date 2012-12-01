// Contains code for all in game entities and objects.
// file pogomud/unierse/universe.go

package world

import (
	"strings"
	"errors"
)

type World struc {
	characters EntityRegistry
	rooms EntityRegistry
}

// NOTE Stubs for stuff right now, will move to seperate files later.
type Entity struct {
	id          int
	name        string
	description string
}

// Simple index and registry for entities until I load this to a
// Database

type EntityRegistry struct {
	idMap map[int]Entity
	nameMap map[string]Entity
	err error
}

func (e *EntityRegistry) regErr() error {
	err := e.err
	e.err = nil
	return err
}

func (self *EntityRegistry) GetById(id int) (Entity, error) {
	var err error
	if entity, ok := self.idMap[id]; ok {
		return entity, err
	}
	return nil, errors.New("id not found.")
}

func (self *EntityRegistry) GetByName(name string) Entity {
	if entity, ok := self.nameMap[strings.ToLower(name)]; ok {
		return entity
	}
	return nil, errors.New("name not found.")
}

func (self *EntityRegistry) Add(e Entity) {
	self.idMap[e.id] = e
	self.nameMap[strings.ToLower(e.name)]
}

func (self *EntityRegistry) Remove(e Entity) {
	delete(self.idMap, e.id)
	delete(self.nameMap, strings.ToLower(e.name))
}