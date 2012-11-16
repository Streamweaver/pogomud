// Unitests for user.go

package server

import (
	"testing"
	)

func Test_validateName(t *testing.T) {
	userList := make(map[string]User)
	userList["TestExisting"] = User{}

	badNames := []string{
		"Not 2 pass", // spaces
		"2_not_pass", // leading number
		"Not_Too_pass_and_then_some_and_some_more", // too long
		"no", // too short
		"TestExisting",
	}

	for name := range badNames {
		pass, msg := validateName(badNames[name], userList)
		if pass {
			t.Errorf("NAME INCORRECTLY PASSED '%s' with %s.", badNames[name], msg)
		}
	}

	goodName := "Now2pass"
	pass, msg := validateName(goodName, userList)
	if !pass {
		t.Errorf("NAME INCORRECTLY PASSED '%s' with %s.", goodName, msg)
	}

}