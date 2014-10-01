// // Represents a server user and needed function.
package server

import (
	"bufio"
	"fmt"
	"net"
	"regexp"
)

const (
	VALIDNAMEMSG = "Something helpful"
	WELCOMEMSG = "Welcome to pogomud"
	)

func HandleUser(conn *net.TCPConn) {
	conn.Write([]byte(WELCOMEMSG + "\n"))
	name := nameSetter(conn)
	conn.Write([]byte("Name set to " + name + " goodbye!\n"))
	conn.Close()
}

// // Checks for valid usernames.
func nameSetter(conn *net.TCPConn) string {
	var name string
	conn.Write([]byte(VALIDNAMEMSG + "\n" + "Enter a name to use: "))
	r := bufio.NewReader(conn)
	for name == "" {
		line, _, err := r.ReadLine()
		if err != nil {
			conn.Write([]byte(fmt.Sprintf("Name Error: %s", err)))
			continue
		}
		valid, msg := validateName(string(line))
		if valid {
			name = string(line)
		} else {
			conn.Write([]byte(fmt.Sprintf("Name Error: %s", msg)))
		}
	}
	return name
}

// Checks for valid name string and returns ture of false.
// TODO add validation.
func validateName(name string) (bool, string) {
	namePtrn, err := regexp.Compile(`^[a-zA-Z][a-zA-z_\d]{3,15}$`)
	if err != nil {
		fmt.Printf("Error in name validation regexp: %s", err)
		// TODO log this
	}
	// That username itself follows basic naming rules.
	if !namePtrn.MatchString(name) {
		return false, "Invalid username. " + VALIDNAMEMSG
	}
	// TODO Check regexp here.
	return true, "Name accepted." // all names valid right now.
}
