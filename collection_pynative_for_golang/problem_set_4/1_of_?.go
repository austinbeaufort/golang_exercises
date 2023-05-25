// Create a list by picking odd-index items from the first list and even index items from the second
package main

import "fmt"

type Lists map[string][]int

func Solve1() {
	l1 := []int{3, 6, 9, 12, 15, 18, 21}
	l2 := []int{4, 8, 12, 16, 20, 24, 28}
	oddVals := getOddIndexNums(l1)
	evenVals := getEvenIndexNums(l2)

	lists := Lists{
		"oddVals":  oddVals,
		"evenVals": evenVals,
		"allVals":  append(oddVals, evenVals...),
	}
	printResults(lists)
}

func printResults(lists Lists) {
	fmt.Println("Odd Index Nums from l1: ")
	fmt.Println(lists["oddVals"])
	fmt.Println("Even Index Nums from l2: ")
	fmt.Println(lists["evenVals"])
	fmt.Println("All lists: ")
	fmt.Println(lists["allVals"])

}

func getOddIndexNums(nums []int) []int {
	odds := []int{}
	for i, num := range nums {
		if isOdd(i) {
			odds = append(odds, num)
		}
	}
	return odds
}

func getEvenIndexNums(nums []int) []int {
	evens := []int{}
	for i, num := range nums {
		if isEven(i) {
			evens = append(evens, num)
		}
	}
	return evens
}

func isOdd(num int) bool {
	if num == 0 {
		return false
	}
	return num%2 != 0
}

func isEven(num int) bool {
	return num%2 == 0
}
