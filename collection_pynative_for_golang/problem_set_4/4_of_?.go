// Create a list such that it shows the element from both lists in a pair.
// expected output: [[2, 4], [3, 9], [4, 16]...etc]
package main

import "fmt"

func Solve4() {
	l1 := [7]int{2, 3, 4, 5, 6, 7, 8}
	l2 := [7]int{4, 9, 16, 25, 36, 49, 64}

	fmt.Println(getResults(l1, l2))
}

func getResults(l1, l2 [7]int) [][]int {
	arrLength := len(l1) - 1

	newArr := [][]int{}
	for i := 0; i <= arrLength; i++ {
		littleArr := []int{l1[i], l2[i]}
		newArr = append(newArr, littleArr)
	}
	return newArr
}
