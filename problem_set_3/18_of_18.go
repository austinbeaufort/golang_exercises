// Replace each special symbol with # in the following string
package main

import (
	"fmt"
	"unicode"
)

func Solve18() {
	sentence := "/*Jon is @developer & musician!!"
	fmt.Println(formatStr(sentence))
}

func formatStr(val string) (newStr string) {
	for _, char := range val {
		if unicode.IsDigit(char) || unicode.IsLetter(char) || unicode.IsSpace(char) {
			newStr += string(char)
		} else {
			newStr += "#"
		}
	}
	return
}
