// Write a program to find value 20 in the list, and if it is present, replace it with 200
// Only update the first occurrence of an item.
package main

import "fmt"

func Solve11() {
	arr := [7]int{5, 10, 15, 20, 25, 50, 20}
	fmt.Println(replaceFirst20Instance(arr))
}

func replaceFirst20Instance(arr [7]int) [7]int {
	for i, num := range arr {
		if num == 20 {
			arr[i] = 200
			return arr
		}
	}
	panic("No value of 20 found in given arr")
}
