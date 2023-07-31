// Return true if the string in the first element of the array contains all of the letters of the string in the second element of the array.
package main

import (
	"fmt"
	"strings"
)

func Solve11() {
	arr1 := [2]string{"hello", "Hello"}
	arr2 := [2]string{"hello", "hey"}
	arr3 := [2]string{"Alien", "line"}
	arr4 := [2]string{"zyxwvutsrqponmlkjihgfedcba", "qrstu"}
	arr5 := [2]string{"Mary", "Aarmy"}

	fmt.Println(containsAllLetters(arr1))
	fmt.Println(containsAllLetters(arr2))
	fmt.Println(containsAllLetters(arr3))
	fmt.Println(containsAllLetters(arr4))
	fmt.Println(containsAllLetters(arr5))
}

func containsAllLetters(arr [2]string) bool {
	lowerWord1 := strings.ToLower(arr[0])
	lowerWord2 := strings.ToLower(arr[1])
	for _, char := range lowerWord2 {
		if !strings.Contains(lowerWord1, string(char)) {
			return false
		}
	}
	return true
}
