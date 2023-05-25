// Accept a list of 5 float numbers as an input from the user, store them as an array and print the array
package main

import "fmt"

func Solve5() {
	fmt.Println("Enter five floats: ")
	fmt.Println(makeArr())
}

func makeArr() [5]float64 {
	arr := [5]float64{}
	for i := 0; i < 5; i++ {
		arr[i] = GetFloat()
	}
	return arr
}

func GetFloat() float64 {
	var num float64
	_, err := fmt.Scan(&num)
	HandleErr(err)
	return num
}
