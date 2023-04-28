// reverse a given string
package main

import "fmt"

func Solve11() {
	fmt.Println("Enter a string to reverse: ")
	val := GetStrInput()
	fmt.Println(revStr2(val))
}

func revStr2(val string) (newVal string) {
	lastIndex := len(val) - 1
	for i := lastIndex; i >= 0; i-- {
		newVal += string(val[i])
	}
	return
}
