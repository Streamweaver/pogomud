// Handles managing all users on the server.
package server 

import (
	"strings"
)

type userRegistry struct {
	byId map[int]User
	byName map[string]User
}

func (self *userRegistry)addUser(user *User) {
	self.idList[u.id] := user
	self.nameList[strings.ToLower(user.name)] := user
}

func (self *userRegistry)removeUser(user *User) {
	delete(self.byId, user.id)
	delete(self.byName, strings.ToLower(user.name))
}