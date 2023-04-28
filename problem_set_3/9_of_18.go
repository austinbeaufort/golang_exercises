// Given a string s1, write a program to return the sum and average of the digits that appear in the string, ignoring all other characters.
package main

import (
	"fmt"
	"unicode"
)

func Solve9() {
	val := "PYnative29@#8496"
	nums := getDigits(val)
	sum := getSum(nums)
	average := float64(sum) / float64(len(nums))
	fmt.Printf("Sum: %d Average: %f\n", sum, average)
}

func getSum(nums []int) (total int) {
	for _, num := range nums {
		total += num
	}
	return
}

func getDigits(val string) []int {
	var digits []int
	for _, char := range val {
		if unicode.IsDigit(char) {
			digits = append(digits, int(char-'0'))
		}
	}
	return digits
}
