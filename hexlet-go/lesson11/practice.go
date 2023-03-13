package lesson11

import (
	"strings"
)

type UserCreateRequest struct {
	FirstName string
	Age       int
}

// BEGIN (write your solution here)
func Validate(req UserCreateRequest) string {
	if isValidName(req.FirstName) && isValidAge(req.Age) {
		return ""
	}
	return "invalid request"
}

func isValidName(name string) bool {
	return !(name == "" || strings.Contains(name, " "))
}

func isValidAge(age int) bool {
	return !(age == 0 || age < 0 || age > 150)
}

// END
