// // Represents a server user and needed function.
package server

import (
	"bufio"
	"fmt"
	"net"
	"regexp"
	"strings"
	"github.com/streamweaver/pogomud/world"
)

const (
	VALIDNAMEMSG = "Something helpful"
	WELCOMEMSG = "Welcome to pogomud"
	)

func HandleUser(conn *net.TCPConn) {
	conn.Write([]byte(WELCOMEMSG + "\n"))
	name := promptName(conn)
	conn.Write([]byte("Name set to " + name + " goodbye!\n"))
	conn.Close()
}

// // Checks for valid usernames.
func promptName(conn *net.TCPConn) string {
	var name string
	conn.Write([]byte(VALIDNAMEMSG + "\n" + "Enter a character name: "))
	r := bufio.NewReader(conn)
	for name == "" {
		line, _, err := r.ReadLine()
		n := strings.TrimSpace(string(line))
		if err != nil {
			conn.Write([]byte(fmt.Sprintf("Name Error: %s", err)))
			continue
		}
		valid, msg := validateName(n)
		if valid {
			name = n
		} else {
			conn.Write([]byte(fmt.Sprintf("Name Error: %s", msg)))
		}
	}
	return name
}

func newCharacter(conn *net.TCPConn) *world.Player {
	name := setName(conn)
	desc := setDescription(conn)
}

func setName(conn *net.TCPConn) string {
	var name string
	conn.Write([]byte(VALIDNAMEMSG + "\n" + "Enter a new character name: "))
	r := bufio.NewReader(conn)
	for name == "" {
		line, _, err := r.ReadLine()
		n := strings.TrimSpace(string(line))
		if err != nil {
			conn.Write([]byte(fmt.Sprintf("Name Error: %s", err)))
			continue
		}
		valid, msg := validateName(n)
		if valid {
			name = n
		} else {
			conn.Write([]byte(fmt.Sprintf("Name Error: %s", msg)))
		}
	}
	return name
}

func setDescription(conn *net.TCPConn) string {
	var desc string
	conn.Write([]byte("Enter a short character description: "))
	r := bufio.NewReader(conn)
	for desc == "" {
		line, _, err := r.ReadLine()
		d := strings.TrimSpace(string(line))
		if err != nil {
			conn.Write([]byte(fmt.Sprintf("Error: %s", err)))
			continue
		}
		valid, msg := validateDescription(d)
		if valid {
			desc = d
		} else {
			conn.Write([]byte(fmt.Sprintf("Error: %s", msg)))
		}
	}
	return desc
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

func validateDescription(desc string) (bool, string) {
	if desc == "" {
		return false, ""
	}
	return true, desc
}
