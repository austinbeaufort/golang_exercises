// We'll pass you an array of two numbers.
// Return the sum of those two numbers plus the sum of all the numbers between them.
// The lowest number will not always come first.
package main

import (
	"fmt"
	"sort"
)

func Solve1() {
	arr1 := []int{4, 1}
	fmt.Println(sumRangeInclusive(arr1))
	arr2 := []int{5, 10}
	fmt.Println(sumRangeInclusive(arr2))
	arr3 := []int{10, 5}
	fmt.Println(sumRangeInclusive(arr3))
}

func sumRangeInclusive(arr []int) (total int) {
	sort.Ints(arr)
	for i := arr[0]; i <= arr[1]; i++ {
		total += i
	}
	return
}
