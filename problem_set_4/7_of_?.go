// Iterate a given list and check if a given element exists as a key’s value in a dictionary. If not, delete it from the list.
package main

import "fmt"

func Solve7() {
	nums := [8]int{47, 64, 69, 37, 76, 83, 95, 97}
	people := map[string]int{
		"John":  47,
		"Emma":  69,
		"Kelly": 76,
		"Jason": 97,
	}
	fmt.Printf("Nums list after removing nums not associated with people: %v\n", getNumsofPeople(nums, people))
}

func getNumsofPeople(nums [8]int, people map[string]int) []int {
	filteredNums := []int{}
	for _, num := range nums {
		if numOfPerson(num, people) {
			filteredNums = append(filteredNums, num)
		}
	}
	return filteredNums
}

func numOfPerson(num int, people map[string]int) bool {
	for _, val := range people {
		if num == val {
			return true
		}
	}
	return false
}
