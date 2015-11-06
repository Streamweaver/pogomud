package pogomud

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

type Server struct {
	Name        string `json:"name"`         // Name of the MUD server.
	Host        string `json:"host"`         // IP or DNS of host.
	Port        int    `json:"port"`         // Port to run on.
	BufferLimit int    `json:"buffer_limit"` // Buffer size limit to use.
	RejectMsg   string `json:"reject_msg"`   // String to send if not accepting connections.
	allowLogin  bool   // Flag to allow connections
	running     bool   // Boolean of weather server should be on or off.
}

// Returns a new Server instance configured with json file provided by cfg
func NewServer() *Server {

	// Load json file with config information.
	file, err := ioutil.ReadFile(cfg)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Parse file info into data
	server := new(Server)
	if err = json.Unmarshal(file, server); err != nil {
		log.Fatal(err)
	}

	return server
}

// Creates a server on the host address and port and opens it for connections.
func (s *Server) Start() error {
	s.running = true
	s.allowLogin = true

	// Setup the server address and listener.
	addr, err := net.ResolveTCPAddr("tcp", net.JoinHostPort(s.Host, fmt.Sprintf("%d", s.Port)))
	if err != nil {
		log.Fatal(err)
		return err
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Printf("%s server started and listening on port %d.\n", s.Name, s.Port)

	// Listen for and accept user connections until server shutdown.
	for s.running {
		// Wait for a connection.
		conn, err := l.AcceptTCP()
		if err != nil {
			log.Fatal(err)
			return err
		}

		if s.allowLogin {
			go HandleUser(conn)
			log.Printf("Connection made from %s\n", conn.RemoteAddr())
		} else {
			s.RejectConnection(conn)
		}
	}

	log.Printf("Server Stopped, exiting.")
	return nil
}

func (s *Server) Stop() {
	log.Printf("Stopping Server.")
	s.allowLogin = false
	s.running = false
	// TODO any other cleanup I need to do.
}

func (s *Server) RejectConnection(conn *net.TCPConn) {
	conn.Write([]byte(s.RejectMsg))
	conn.Write([]byte("\n"))
	log.Printf("Server refused connection from %s", conn.RemoteAddr())
	conn.Close()
}
