// Iterate the given list of numbers and print only those numbers which are divisible by 5
package main

import "fmt"

func main() {
	arr := [5]int{10, 20, 33, 46, 5000}
	fmt.Println("Array values divisible by 5: ")
	printDivBy5(arr)
}

func printDivBy5(arr [5]int) {
	for _, num := range arr {
		if num%5 == 0 {
			fmt.Println(num)
		}
	}
}
