// file gomud/gomud.go - Entry point for PoGoMud
package main

import (
	"github.com/streamweaver/server"
)

func main() {
	server := server.NewServer()
	server.Start("config.json") // TODO make this callable by flag.
}
