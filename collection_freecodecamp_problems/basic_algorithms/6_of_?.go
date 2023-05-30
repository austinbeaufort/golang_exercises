// Truncate a string (first argument) if it is longer than the given maximum string length (second argument).
// Return the truncated string with a ... ending.
package main

import "fmt"

func Solve6() {
	fmt.Println("Enter a sentence to truncate: ")
	val := getStringInput()
	fmt.Println("Enter an index (number) at which to truncate the given sentence: ")
	num := getInt()
	fmt.Println(truncateString(val, num))
}

func truncateString(val string, num int) string {
	if len(val) - 1 <= num {
		return val
	}

	return val[:num] + "..."
}