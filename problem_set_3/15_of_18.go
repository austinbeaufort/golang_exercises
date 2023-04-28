// Remove special symbols / punctuation from a string
package main

import (
	"fmt"
	"unicode"
)

func Solve15() {
	sentence := "/*Jon is @developer & musician."
	fmt.Println(removeSymbols(sentence))
}

func removeSymbols(sentence string) (filteredVal string) {
	for _, char := range sentence {
		if unicode.IsDigit(char) || unicode.IsLetter(char) || unicode.IsSpace(char) {
			filteredVal += string(char)
		}
	}
	return
}
