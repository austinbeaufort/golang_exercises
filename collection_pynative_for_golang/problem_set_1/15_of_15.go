// Write a function called exponent(base, exp) that returns an int value of base raises to the power of exp.
// We are assuming no built in function allowed for this exercise.
package main

import "fmt"

func main() {
	num1, num2 := getInput()
	fmt.Printf("Base: %d to the power: %d = %d\n", num1, num2, pow(num1, num2))
}

func pow(base, exponent int) int {
	total := 1
	for i := 1; i <= exponent; i++ {
		total *= base
	}
	return total
}

func getInput() (int, int) {
	fmt.Println("Enter a Base Value and an Exponent: ")
	var num1, num2 int
	_, err := fmt.Scan(&num1, &num2)
	handleErr(err)
	return num1, num2
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
