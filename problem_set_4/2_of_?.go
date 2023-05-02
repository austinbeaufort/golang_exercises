// Separate list of nine ints into 3 equal chunks and reverse each chunk
package main

import "fmt"

func Solve2() {
	list := [9]int{11, 45, 8, 23, 14, 12, 78, 45, 89}
	lists := map[string][]int{
		"chunk1": list[0:3],
		"chunk2": list[3:6],
		"chunk3": list[6:],
	}
	printLists(lists)
}

func printLists(lists map[string][]int) {
	for key, value := range lists {
		fmt.Printf("%s --> %v\n", key, value)
		fmt.Printf("After reversing: %v\n", revList(value))
	}
}

func revList(list []int) []int {
	newList := []int{}
	lastIndex := len(list) - 1

	for i := lastIndex; i >= 0; i-- {
		newList = append(newList, list[i])
	}
	return newList
}
