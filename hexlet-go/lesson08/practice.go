package lesson08

import "strings"

// BEGIN (write your solution here)
func Greetings(name string) string {
	return strings.Title("привет, " + strings.ToLower(strings.Trim(name, " ")) + "!")
}

// END
