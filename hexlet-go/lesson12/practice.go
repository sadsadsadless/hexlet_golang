package lesson12

// BEGIN (write your solution here)
func ErrorMessageToCode(msg string) int {
	const (
		OK        = 0
		CANCELLED = 1
		UNKNOWN   = 2
	)
	switch msg {
	case "OK":
		return OK
	case "CANCELLED":
		return CANCELLED
	default:
		return UNKNOWN
	}
}

// END
