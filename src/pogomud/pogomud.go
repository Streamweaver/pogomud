// file gomud/gomud.go - Entry point for PoGoMud
package main

import (
	"pogomud/server"
)

func main() {
	server := server.NewServer()
	server.Start()
}