package helper

import "strings"

const AppName = "Go Basics Application"

func ValidEmail(email string) (bool, bool) {
	return len(email) > 2 && strings.Contains(email, "@"), true
}
