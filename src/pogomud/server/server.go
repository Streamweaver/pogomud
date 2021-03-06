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

type ServerObject struct {
	Name        string       // Name of the MUD server.
	Protocol    string       // Type of TCP protocol to use.
	Host        string       // IP or DNS of host.
	Port        int          // Port to run on.
	BufferLimit int          // Buffer size limit to use.
	Database    DatabaseInfo // Database connection information.
}

type DatabaseInfo struct {
	Host   string // Hostname of Database
	Port   string // Port number of Database
	Name   string // Name of Database
	User   string // Admin Username for Database
	Pass   string // Password for User above.
	Engine string // Type of Database being connected to.
}

// Loads config information from JSON file into a Server Object
func NewServer() ServerObject {

	// Load json file with config information.
	file, e := ioutil.ReadFile("config.json")
	if e != nil {
		fmt.Printf("%s\n", e)
		log.Fatal(e)
	}

	// Parse file info into data
	var server ServerObject
	e = json.Unmarshal(file, &server)
	if e != nil {
		log.Fatal(e)
		os.Exit(1)
	}

	return server
}

// Creates a server on the host address and port and opens it for connections.
func (server *ServerObject) Start() {

	// Setup the server address and listener.
	addr, err := net.ResolveTCPAddr("tcp", net.JoinHostPort(server.Host, fmt.Sprintf("%d", server.Port)))
	if err != nil {
		log.Fatal(err)
		return
	}
	l, err := net.ListenTCP(server.Protocol, addr)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Setup stuff to handle user communication.
	userList := make(map[string]User)
	msgQueue := make(chan Message)

	// Startup a Command Handler to interpret user input.
	// go commandHandler(msgQueue, userList)

	log.Printf("%s server started and listening on port %d.\n", server.Name, server.Port)

	// Listen for and accept user connections.
	for {
		// Wait for a connection. 
		conn, err := l.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}
		// More code here for what to do.
		HandleUser(conn, msgQueue, userList)
		log.Printf("Connection made from %s\n", conn.RemoteAddr())
	}
}
