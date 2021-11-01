package utils

func TrimToDigit(str string) int {
	return 0
}

// rebrain tests
func InSlice(a []string, x string) bool {
	for _, n := range a { //
		if x == n {
			return true
		}
	}
	return false
}

// rebrain tests
func InSliceInt(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
