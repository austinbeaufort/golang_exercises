// Find the intersection (common) of two arrays and remove those elements from the first array
package main

import "fmt"

func Solve5() {
	l1 := [7]int{23, 42, 65, 57, 78, 83, 29}
	l2 := [7]int{57, 83, 29, 67, 73, 43, 48}
	newList1 := removeCommonsFromL1(l1, l2)
	fmt.Println(newList1)
}

func removeCommonsFromL1(l1, l2 [7]int) []int {
	newL1 := []int{}
	for _, val := range l1 {
		if notInL2(val, l2) {
			newL1 = append(newL1, val)
		}
	}
	return newL1
}

func notInL2(val int, l2 [7]int) bool {
	for _, num := range l2 {
		if val == num {
			return false
		}
	}
	return true
}
