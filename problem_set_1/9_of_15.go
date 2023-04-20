// Write a program to check if the given number is a palindrome number.
package main

import "fmt"

func main() {
	val := getInput()
	reversedVal := reverseString(val)
	fmt.Println(checkPalindrome(val, reversedVal))

}

func checkPalindrome(val1, val2 string) string {
	if val1 == val2 {
		return "Your string is a palindrome"
	}
	return "Not a palindrome"
}

func reverseString(val string) string {
	revStr := ""
	last_index := len(val) - 1
	for i := last_index; i >= 0; i-- {
		revStr += string(val[i])
	}
	return revStr
}

func getInput() string {
	fmt.Println("Enter a string to check if it's a palindrome: ")
	var val string
	_, err := fmt.Scan(&val)
	handleErr(err)
	return val
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
