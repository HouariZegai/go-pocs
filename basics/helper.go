package main

import "strings"

func validEmail(email string) (bool, bool) {
	return len(email) > 2 && strings.Contains(email, "@"), true
}
