// Stuff for handling the server itself.
// Playing around with example from https://gist.github.com/775526
package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

type Server struct {
	Name        string // Name of the MUD server.
	Host        string // IP or DNS of host.
	Port        int    // Port to run on.
	BufferLimit int    // Buffer size limit to use.
}

// Loads config information from JSON file provided in the pth argument.
func NewServer(pth string) *Server {

	// Load json file with config information.
	file, err := ioutil.ReadFile()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Parse file info into data
	server := new(Server)
	e = json.Unmarshal(file, server)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return server
}

// Creates a server on the host address and port and opens it for connections.
func (s *Server) Start() {

	// Setup the server address and listener.
	addr, err := net.ResolveTCPAddr("tcp", net.JoinHostPort(s.Host, fmt.Sprintf("%d", s.Port)))
	if err != nil {
		log.Fatal(err)
		return
	}
	l, err := net.ListenTCP(server.Protocol, addr)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("%s server started and listening on port %d.\n", s.Name, s.Port)

	// Listen for and accept user connections.
	for {
		// Wait for a connection.
		conn, err := l.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}
		// More code here for what to do.
		HandleUser(conn)
		log.Printf("Connection made from %s\n", conn.RemoteAddr())
	}
}
