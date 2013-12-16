// Handles things related to in world player characters.
package world

type Character struct {
	Entity
	area    Area
	message chan string
}
