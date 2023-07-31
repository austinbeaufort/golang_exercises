// Return the lowest index at which a value (second argument) should be inserted into an array (first argument) once it has been sorted.
// The returned value should be a number.
package main

import (
	"fmt"
	"sort"
)

func Solve10() {
	vals1 := []int{20, 3, 5}
	num1 := 19

	vals2 := []int{10, 20, 30, 40, 50}
	num2 := 35

	vals3 := []int{40, 60}
	num3 := 50

	vals4 := []int{3, 10, 5}
	num4 := 3

	vals5 := []int{}
	num5 := 1

	fmt.Println(getIndex(vals1, num1))
	fmt.Println(getIndex(vals2, num2))
	fmt.Println(getIndex(vals3, num3))
	fmt.Println(getIndex(vals4, num4))
	fmt.Println(getIndex(vals5, num5))
}

func getIndex(vals []int, num int) int {
	if len(vals) == 0 {
		return 0
	}
	sort.Ints(vals)
	for i, val := range vals {
		if num <= val {
			return i
		}
	}
	panic("index invalid, closing program")
}
