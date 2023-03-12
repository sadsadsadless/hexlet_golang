package lesson06

import "math"

// BEGIN (write your solution here)
func MinInt(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

// END
