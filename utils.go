package utils

// rebrain tests
func InSlice(a []string, x string) bool {
	for _, n := range a { //
		if x == n {
			return true
		}
	}
	return false
}

<<<<<<< HEAD
// rebrain tests
=======
func InSliceInt(a []int, x int) bool {
>>>>>>> master
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
