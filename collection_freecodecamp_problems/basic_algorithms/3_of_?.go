// Return an array consisting of the largest number from each provided sub-array.
// For simplicity, the provided array will contain exactly 4 sub-arrays.
package main

import "fmt"

func Solve3() {
	arrs := [4][4]int {
		{4, 5, 1, 3},
		{13, 27, 18, 26},
		{32, 35, 37, 39},
		{1000, 1001, 857, 1},
	}

	arrs2:= [4][4]int {
		{17, 23, 25, 12},
		{25, 7, 34, 48},
		{4, -10, 18, 21},
		{-72, -3, -17, -10},
	}


	largestNums := getLargestNums(arrs)
	largestNums2 := getLargestNums((arrs2))
	fmt.Println(largestNums)
	fmt.Println(largestNums2)
}

func getLargestNums(arrs [4][4]int) []int {
	largestVals := []int{}
	for _, arr := range arrs {
		largestVal := getLargestVal(arr)
		largestVals = append(largestVals, largestVal)
	} 
	return largestVals
}

func getLargestVal(arr [4]int) (largestNum int) {
	for i, num := range arr {
		if i == 0 || num > largestNum{
			largestNum = num
		}
	}
	return	
}