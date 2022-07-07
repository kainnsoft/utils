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

// if slice of runes is palindrome returns TRUE, else returns FALSE
func IsPalindrome(runes []rune) bool {
	lenth := len(runes)
	for i := 0; i < lenth/2; i++ {
		if runes[i] != runes[lenth-i-1] {
			return false
		}
	}
	return true
}

// Reverse strings, including such:
// `Hello ыыы
//   world! @[ъ]`
func reverceString(input string) (output string) {
	runeInput := []rune(input)
	runeLen := len(runeInput)
	outputRune := make([]rune, runeLen)

	for k, v := range runeInput {
		outputRune[runeLen-k-1] = v
	}
	output = string(outputRune)
	return
}

///////////////////////////////////////////////////
// AsyncFib
func AsyncFib_main_analogue() {
	c := make(chan int, 3)
	quit := make(chan struct{})
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- struct{}{}
	}()
	asyncFib(c, quit)
}

func asyncFib(c chan int, q chan struct{}) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-q:
			fmt.Println("Quit")
			return
		}
	}
}
