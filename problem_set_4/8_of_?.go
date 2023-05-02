// Get all values from the map and add them to a list but don’t add duplicates
package main

import "fmt"

func Solve8() {
	monthsAndNums := map[string]int{
		"jan":   47,
		"feb":   52,
		"march": 47,
		"April": 44,
		"May":   52,
		"June":  53,
		"july":  54,
		"Aug":   44,
		"Sept":  54,
	}
	fmt.Println(createNonDuplicateNumsList(monthsAndNums))
}

func createNonDuplicateNumsList(monthsAndNums map[string]int) []int {
	nums := []int{}
	for _, num := range monthsAndNums {
		if !contains3(num, nums) {
			nums = append(nums, num)
		}
	}
	return nums
}

func contains3(num int, nums []int) bool {
	for _, val := range nums {
		if num == val {
			return true
		}
	}
	return false
}
