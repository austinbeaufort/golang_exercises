// Write a program to accept two numbers from the user and calculate multiplication
package main

import "fmt"

func Solve1() {
	fmt.Println("Enter two numbers to multiple them together: ")
	num1, num2 := GetTwoInts()
	fmt.Println(num1 * num2)
}
