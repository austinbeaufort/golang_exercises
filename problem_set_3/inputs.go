package main

import "fmt"

func GetStrInput() (val string) {
	_, err := fmt.Scan(&val)
	Check(err)
	return
}

func GetTwoStrings() (val1, val2 string) {
	_, err := fmt.Scan(&val1, &val2)
	Check(err)
	return
}

func GetValidInput() string {
	for {
		val := GetStrInput()
		if checkValid(val) {
			return val
		}
		fmt.Println("String entered was invalid, please enter a valid String: ")
	}
}

func checkValid(word string) bool {
	oddNumOfChars := len(word)%2 != 0
	properLength := len(word) >= 3
	return oddNumOfChars && properLength
}
