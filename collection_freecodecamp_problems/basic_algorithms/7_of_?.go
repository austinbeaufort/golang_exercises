// Create a function that looks through an array arr and returns the first element in it that passes a 'truth test'.
// This means that given an element x, the 'truth test' is passed if func(x) is true. If no element passes the test, return undefined.
// also learning syntax for writing functions in functions with golang. (this example is not an example of closure, but nested functions could be used for that concept)

package main

import "fmt"

func Solve7() {
	arr := [4]int {1, 2, 3, 4}
	arr2 := [4]int {1, 3, 5, 7}

	var isEven = func(num int) bool {
		return num % 2 == 0
	}

	fmt.Println(findElements(arr, isEven))
	fmt.Println(findElements(arr2, isEven))
}

func findElements(arr [4]int, fn func(num int) bool) bool {
	fmt.Println(arr)
	for _, num := range arr {
		if fn(num) {
			return true
		}
	}
	return false
}