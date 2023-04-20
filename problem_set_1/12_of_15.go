// Calculate income tax for the given income by adhering to the below rules
// first 10k , rate: 0%
// next 10k  , rate: 10%
// remaining, rate: 20%

package main

import "fmt"

func main() {
	income := getInput()
	fmt.Println(calculateTax(income))
}

func calculateTax(income int) float64 {
	if income <= 10_000 {
		return 0
	}
	if income <= 20_000 {
		extra_income := income - 10_000
		return float64(extra_income) * 0.1
	}
	low_tax := 10_000 * 0.1
	high_tax := float64(income-20_000) * 0.2
	return low_tax + high_tax
}

func getInput() int {
	fmt.Print("Enter your Income: ")
	var income int
	_, err := fmt.Scan(&income)
	handleErr(err)
	return income
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
