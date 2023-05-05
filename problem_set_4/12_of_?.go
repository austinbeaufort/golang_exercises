// Below are two lists.
// Write a program to convert them into a dictionary in a way that item from list1 is the key and item from list2 is the value.
package main

import "fmt"

func Solve12() {
	l1 := [3]string{"Ten", "Twenty", "Thirty"}
	l2 := [3]int{10, 20, 30}
	fmt.Println(makeMap(l1, l2))
}

func makeMap(l1 [3]string, l2 [3]int) map[string]int {
	newMap := make(map[string]int)
	for i, val := range l1 {
		newMap[val] = l2[i]
	}
	return newMap
}
