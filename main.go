// file gomud/gomud.go - Entry point for PoGoMud
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/streamweaver/pogomud/server"
)

func main() {

	var fName string
	flag.StringVar(&fName, "conf", "config.json", "path to config file.")

	// Load json file with config information.
	file, err := ioutil.ReadFile(fName)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Parse file info into data
	server := server.NewServer()
	if err = json.Unmarshal(file, server); err != nil {
		log.Fatal(err)
	}

	go server.Start()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "exit" {
			server.running = false
		}
	}
}
