// Write a Program to extract each digit from an integer in the reverse order.
// For example, If the given int is 7536, the output shall be “6 3 5 7“, with a space separating the digits.
package main

import "fmt"

func main() {
	input := fmt.Sprint(getInput())
	fmt.Println(format(input))
}

func format(val string) string {
	newVal := ""
	lastIndex := len(val) - 1
	for i := lastIndex; i >= 0; i-- {
		newVal += string(val[i]) + " "
	}
	return newVal
}

func getInput() int {
	fmt.Print("Enter a number to format: ")
	var num int
	_, err := fmt.Scan(&num)
	handleErr(err)
	return num
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
