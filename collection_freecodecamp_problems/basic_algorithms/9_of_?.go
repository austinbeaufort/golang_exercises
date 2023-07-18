// You are given two arrays and an index.
// Copy each element of the first array into the second array, in order.
// Begin inserting elements at index n of the second array.
// Return the resulting array. The input arrays should remain the same after the function runs.
package main

import "fmt"

type StrOrInt interface {
	string | int
}

func Solve9() {
	// scenario 1
	arrs1 := map[string][]int {
		"1": {1,2,3},
		"2": {4,5},
	}
	index1 := 1
	fmt.Println(spliceWithRules(arrs1, index1))

	// scenario 2
	arrs2 := map[string][]int {
		"1": {1,2,3,4},
		"2": {},
	}
	index2 := 0
	fmt.Println(spliceWithRules(arrs2, index2))

	// scenario 3
	arrs3 := map[string][]string {
		"1": {"hamster", "turtle"},
		"2": {"cat", "dog", "bird", "fish"},
	}
	index3 := 2
	fmt.Println(spliceWithRulesStr(arrs3, index3))
}


// chose to use without generics for this example.
func spliceWithRules(arrs map[string][]int, index int) []int {
	startArr2, endArr2 := splitArr2(arrs["2"], index)
	formattedArr := []int{}
	formattedArr = append(formattedArr, startArr2...)
	formattedArr = append(formattedArr, arrs["1"]...)
	formattedArr = append(formattedArr, endArr2...)
	return formattedArr
}

func spliceWithRulesStr(arrs map[string][]string, index int) []string {
	startArr2, endArr2 := splitArr2Str(arrs["2"], index)
	formattedArr := []string{}
	formattedArr = append(formattedArr, startArr2...)
	formattedArr = append(formattedArr, arrs["1"]...)
	formattedArr = append(formattedArr, endArr2...)
	return formattedArr
}

func splitArr2(arr []int, index int) ([]int, []int) {
	return arr[:index], arr[index:] 
}

func splitArr2Str(arr []string, index int) ([]string, []string) {
	return arr[:index], arr[index:] 
}