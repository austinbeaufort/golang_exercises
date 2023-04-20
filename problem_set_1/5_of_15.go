// Write a function to return True if the first and last number of a given list is same. If numbers are different then return False.
package main

import "fmt"

func main() {
	arr1 := [5]int{10, 20, 30, 40, 10}
	arr2 := [5]int{75, 65, 35, 75, 30}
	res1 := checkFirstLast(arr1)
	res2 := checkFirstLast(arr2)
	fmt.Println(res1)
	fmt.Println(res2)
}

func checkFirstLast(arr [5]int) bool {
	lastIndex := len(arr) - 1
	return arr[0] == arr[lastIndex]
}
