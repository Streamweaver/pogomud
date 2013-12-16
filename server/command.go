// Basic file to handle MUD Commands.
package server

import (
	"fmt"
)

func CommandParser(u *User, line string) {
	// TODO add actual parsing here.  for now everyting is a say.
	say(u, line)
}

type command struct {
	name    string
	help    string
	execute func(actor *User, line string)
}

func NewCommand(name string, help string, f func(*User, string)) *Command {
	cmd := &Command{
		name:    name,
		help:    help,
		execute: f,
	}
	return cmd
}
