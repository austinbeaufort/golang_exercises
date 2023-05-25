package main

import "fmt"

func GetInt() int {
	var num int
	_, err := fmt.Scan(&num)
	HandleErr(err)
	return num
}

func GetTwoInts() (int, int) {
	var num1, num2 int
	_, err := fmt.Scan(&num1, &num2)
	HandleErr(err)
	return num1, num2
}

func GetString() string {
	var val string
	_, err := fmt.Scan(&val)
	HandleErr(err)
	return val
}
