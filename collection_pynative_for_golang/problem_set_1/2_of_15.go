// Write a program to iterate the first 10 numbers and in each iteration, print the sum of the current and previous number.

package main

import "fmt"

func main() {
	previous_num := 0
	for i := 0; i < 10; i++ {
		fmt.Printf("Current Number %d Previous Number %d Sum: %d\n", i, previous_num, i+previous_num)
		previous_num = i
	}
}
