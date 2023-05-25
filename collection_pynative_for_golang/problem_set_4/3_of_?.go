// Write a program to iterate a given list and count the occurrence
// of each element and create a map to show the count of each element.
package main

import "fmt"

type Counts map[int]int

func Solve3() {
	nums := [9]int{11, 45, 8, 11, 23, 45, 23, 45, 89}
	uniques := getUniques(nums)
	fmt.Println(getCounts(uniques, nums))

}

func getCounts(uniques []int, nums [9]int) Counts {
	counts := make(Counts)
	for _, num := range uniques {
		count := countNum(num, nums)
		counts[num] = count
	}
	return counts
}

func countNum(num int, nums [9]int) (count int) {
	for _, val := range nums {
		if val == num {
			count += 1
		}
	}
	return
}

func getUniques(nums [9]int) []int {
	uniques := []int{}
	for _, num := range nums {
		if !contains(num, uniques) {
			uniques = append(uniques, num)
		}
	}
	return uniques
}

func contains(val int, nums []int) bool {
	for _, num := range nums {
		if val == num {
			return true
		}
	}
	return false
}
