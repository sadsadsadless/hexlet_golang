package lesson10

import "strings"

// BEGIN (write your solution here)
func ModifySpaces(s, mode string) string {
	switch mode {
	case "dash":
		return strings.ReplaceAll(s, " ", "-")
	case "underscore":
		return strings.ReplaceAll(s, " ", "_")
	default:
		return strings.ReplaceAll(s, " ", "*")
	}
}

// END
