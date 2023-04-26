// Display three string “Name”, “Is”, “James” as “Name**Is**James”
package main

import "fmt"

func Solve2() {
	fmt.Println("Enter three strings to format: ")
	val1, val2, val3 := getStrings()
	fmt.Printf("%s**%s**%s\n", val1, val2, val3)
}

func getStrings() (string, string, string) {
	var val1, val2, val3 string
	_, err := fmt.Scan(&val1, &val2, &val3)
	HandleErr(err)
	return val1, val2, val3
}
