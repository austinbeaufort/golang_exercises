// Remove all characters from a string except for integers
package main

import (
	"fmt"
	"unicode"
)

func Solve16() {
	sentence := "I am 25 years and 10 months old"
	fmt.Println(numsOnly(sentence))
}

func numsOnly(val string) (newVal string) {
	for _, char := range val {
		if unicode.IsDigit(char) {
			newVal += string(char)
		}
	}
	return
}
