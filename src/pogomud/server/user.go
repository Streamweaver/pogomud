// Represents a MUD user and needed function.
package server

import (
	"net"
	"log"
	"bufio"
	"fmt"
)

type User struct {
	Name string
	Conn *net.TCPConn
	toServer chan Message
	toUser chan Message
}

// Closes the connection and preforms anything needed with it.
func (u *User) Close() {
	//
}

// Removes the user from the server userlist.
func (u *User) Destroy() {
	//
}

// Listends to the users Outgoing channel and sends
// new values to the connection.
func HandleToServer(user *User) {
	reader := bufio.NewReader(user.Conn)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		user.toServer <- NewMessage(user.Name, string(line))
	}
}

// Reads their connection buffer and sends to message.
func HandleToUser(user *User) {
	for {
		msg := <- user.toUser
		text := "(" + msg.name + "): " + msg.content + "\n"
		user.Conn.Write([]byte(text))
	}
}

func nameSetter(conn *net.TCPConn) string {
	conn.Write([]byte("Enter a name to use: "))
	r := bufio.NewReader(conn)
	line, _, err := r.ReadLine()
	if err != nil {
		fmt.Printf("Error reading name string: %s", err)
		log.Fatal(err)
	}
	
	return string(line)
}

// type User struct {
// 	Name string
// 	Conn *net.TCPConn
// 	toServer chan server.Message
// 	toUser chan server.Message
// }
func HandleUser(conn *net.TCPConn, toServer chan Message, userList map[string]User) {
	name := nameSetter(conn)
	newUser := User{
		name,
		conn,
		toServer,
		make(chan Message),
	}
	userList[newUser.Name] = newUser
	go HandleToServer(&newUser)
	go HandleToUser(&newUser)
	toServer <- NewMessage("server", name + " has connected.")
}