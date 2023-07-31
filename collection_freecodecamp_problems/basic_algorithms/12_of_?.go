// Write a function that splits an array (first argument) into groups the length of size (second argument) and returns them as a two-dimensional array.
package main

import "fmt"

func Solve12() {
	arr1 := []int{1, 2, 3, 4}
	size1 := 2

	arr2 := []int{0, 1, 2, 3, 4, 5}
	size2 := 3

	arr3 := []int{0, 1, 2, 3, 4, 5}
	size3 := 2

	arr4 := []int{0, 1, 2, 3, 4, 5}
	size4 := 4

	arr5 := []int{0, 1, 2, 3, 4, 5, 6}
	size5 := 3

	arr6 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	size6 := 2

	fmt.Println(splitIntoSubArrays(arr1, size1))
	fmt.Println(splitIntoSubArrays(arr2, size2))
	fmt.Println(splitIntoSubArrays(arr3, size3))
	fmt.Println(splitIntoSubArrays(arr4, size4))
	fmt.Println(splitIntoSubArrays(arr5, size5))
	fmt.Println(splitIntoSubArrays(arr6, size6))
}

func splitIntoSubArrays(arr []int, size int) [][]int {
	newArr := [][]int{}
	for i := range arr {
		if i != 0 && (i+1)%size == 0 {
			newArr = append(newArr, arr[i-size+1:i+1])
		}
	}
	formattedArr := gatherAndAppendRemainder(newArr, arr)
	return formattedArr
}

func gatherAndAppendRemainder(newArr [][]int, arr []int) [][]int {
	remainderValues := []int{}
	for _, num := range arr {
		isInNewArray := false
		for _, array := range newArr {
			if contains(array, num) {
				isInNewArray = true
			}
		}
		if !isInNewArray {
			remainderValues = append(remainderValues, num)
		}
	}

	if len(remainderValues) > 0 {
		newArr = append(newArr, remainderValues)
	}
	return newArr
}

func contains(arr []int, val int) bool {
	for _, item := range arr {
		if item == val {
			return true
		}
	}
	return false
}
