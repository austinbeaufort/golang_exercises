// Given a two list of numbers,
// write a program to create a new list such that the new list should contain odd numbers from the first list and even numbers from the second list.
package main

import "fmt"

type Options struct {
	Odd, Even string
}

func main() {
	options := Options{"odd", "even"}
	arr1 := [5]int{10, 20, 25, 30, 35}
	arr2 := [5]int{40, 45, 60, 75, 90}
	odds := getOddsOrEvens(arr1, options.Odd)
	evens := getOddsOrEvens(arr2, options.Even)
	result := append(odds, evens...)
	fmt.Println(result)
}

func getOddsOrEvens(arr [5]int, val string) []int {
	newArr := make([]int, 0)
	for _, num := range arr {
		if val == "odd" && num%2 != 0 {
			newArr = append(newArr, num)
		}
		if val == "even" && num%2 == 0 {
			newArr = append(newArr, num)
		}
	}
	return newArr
}
