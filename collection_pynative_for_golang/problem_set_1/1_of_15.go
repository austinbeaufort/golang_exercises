// Given two integer numbers return their product only if the product is equal to or lower than 1000, else return their sum.
package main

import "fmt"

func main() {
	num1, num2 := getNums()
	fmt.Println(getResult(num1, num2))
}

func getResult(num1, num2 int) string {
	if num1*num2 <= 1000 {
		return fmt.Sprintf("Product is %d", num1*num2)
	}
	return fmt.Sprintf("Sum is %d", num1+num2)
}

func getNums() (int, int) {
	fmt.Println("Enter Two Numbers: ")
	var num1, num2 int
	_, err := fmt.Scan(&num1, &num2)
	checkErr(err)
	return num1, num2
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
