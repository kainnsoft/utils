package utils

import (
	"fmt"
	"strconv"
	"unicode"
)

func TrimToDigit(str string) (int, error) {
	var resultStr string
	for _, v := range str {
		if unicode.IsDigit(v) {
			resultStr += string(v)
		}
	}
	retInt, err := strconv.Atoi(resultStr)
	if err != nil {
		return 0, fmt.Errorf("convertation \"%s\" to int failed due to error: %v", str, err)
	}
	return retInt, nil
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
