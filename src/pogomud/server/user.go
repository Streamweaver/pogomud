// Represents a MUD user and needed function.
package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"pogomod/world"
	"regexp"
	"strings"
)

const (
	VALIDNAMEMSG = "Usernames must begin with a letter and contain only letters, understores or numbers."
	WELCOMEMSG   = "Welcome to PoGoMud!"
)

// Represents all the connection and channel information
// needed for a user on the server.
type User struct {
	id        int
	name      string
	conn      *net.TCPConn
	toUser    chan string
	online    bool // Setting to false ends users handlers.
	users     UserRegistry
	room      Room
	character Character
}

// Closes the connection and preforms anything needed with it.
func (self *User) Logout() {
	self.users.removeUser(self)
	self.online = false
	self.conn.Close()
}

// Listends to the users Outgoing channel and sends
// new values to the connection.
func HandleCommands(user *User) {
	reader := bufio.NewReader(user.Conn)
	for user.online {
		bLine, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
			continue
		}
		CommandParser(user, strings.Trim(string(bLine), " "))
	}
}

// Reads their connection buffer and sends to message.
func HandleOut(user *User) {
	for user.online {
		msg := <-user.toUser
		text := "(" + msg.name + "): " + msg.content + "\n"
		user.Conn.Write([]byte(text))
	}
}

func HandleUser(conn *net.TCPConn, toServer chan string, userList map[string]User) {
	conn.Write([]byte(WELCOMEMSG + "\n"))
	name := nameSetter(conn, userList)
	newUser := User{
		name,
		conn,
		toServer,
		make(chan string),
		true,
		userList,
	}
	userList[newUser.Name] = newUser
	go HandleCommands(&newUser)
	go HandleToUser(&newUser)
	toServer <- NewMessage("server", name+" has connected.")
}

// Checks for valid usernames.
func nameSetter(conn *net.TCPConn, UserList map[string]User) string {
	var name string
	for name == "" {
		conn.Write([]byte(VALIDNAMEMSG + "\n" + "Enter a name to use: "))
		r := bufio.NewReader(conn)
		line, _, err := r.ReadLine()
		if err != nil {
			conn.Write([]byte(fmt.Sprintf("Name Error: %s", err)))
			continue
		}
		valid, msg := validateName(string(line), UserList)
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
func validateName(name string, UserList map[string]User) (bool, string) {
	if _, ok := UserList[name]; ok {
		return false, "User already exists with that name."
	}
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
