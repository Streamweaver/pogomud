// file gomud/gomud.go - Entry point for PoGoMud
package main

import (
	"flag"
	"github.com/streamweaver/pogomud/server"
)

func main() {

	var fName string
	flag.StringVar(&fName, "conf", "config.json", "path to config file.")

	server := server.NewServer(fName)
	server.Start()
}
