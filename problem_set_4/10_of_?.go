// Write a program to add item 7000 after 6000 in the nested structure
package main

import "fmt"

type Inner struct {
	num1 int
	num2 int
	arr  []int
	num3 int
}

type Outer struct {
	num1  int
	num2  int
	inner Inner
	num3  int
	num4  int
}

func Solve10() {
	nums := Outer{
		num1: 10,
		num2: 20,
		inner: Inner{
			num1: 300,
			num2: 400,
			arr:  []int{5000, 6000},
			num3: 500,
		},
		num3: 30,
		num4: 40,
	}

	nums.inner.arr = append(nums.inner.arr, 7000)
	fmt.Printf("New Array: %v\n", nums.inner.arr)
	fmt.Println(nums)
}
