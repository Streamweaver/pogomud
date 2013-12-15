// Basic file to handle MUD Commands.
package server

func CommandParser(user *User, line string) {
	// TODO add actual parsing here.  for now everyting is a say.
	say(user, line)
}

func say(speaker *User, line string) {
	for name, user := range speaker.userList {
		if name != speaker.Name {
			user.toUser <- Message{speaker.Name, line}
		}
	}
}
