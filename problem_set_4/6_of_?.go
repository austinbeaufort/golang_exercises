// Checks if one set is a subset or superset of another set.
package main

import "fmt"

func Solve6() {
	l1 := []int{27, 43, 34}
	l2 := []int{34, 93, 22, 27, 43, 53, 48}

	fmt.Printf("First set is subset of second set: %t\n", isSubset(l1, l2))
	fmt.Printf("Second set is subset of first set: %t\n", isSubset(l2, l1))
	fmt.Printf("First set is superset of second set: %t\n", !isSubset(l1, l2))
	fmt.Printf("Second set is superset of first set: %t\n", !isSubset(l2, l1))
}

func contains2(val int, list []int) bool {
	for _, num := range list {
		if val == num {
			return true
		}
	}
	return false
}

func isSubset(l1 []int, l2 []int) bool {
	for _, val := range l1 {
		if !contains2(val, l2) {
			return false
		}
	}
	return true
}
