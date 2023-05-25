// Write a program to use string.format() method to format the following three variables as per the expected output
// Expected Output: I have 1000 dollars so I can buy 3 microwaves for 450.00 dollars.
package main

import "fmt"

func Solve8() {
	money := 1000
	quantity := 3
	var totalSpending float64 = 450

	formattedSpending := fmt.Sprintf("%.2f", totalSpending)
	fmt.Printf("I have %d so I can buy %d microwaves for %s dollars.\n", money, quantity, formattedSpending)
}
