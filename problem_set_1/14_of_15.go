// Print downward Half-Pyramid Pattern with Star
// * * * * *
// * * * *
// * * *
// * *
// *
package main

import "fmt"

func main() {
	for i := 5; i >= 1; i-- {
		fmt.Println(getRow(i))
	}
}

func getRow(num int) string {
	val := ""
	for i := num; i >= 1; i-- {
		val += "* "
	}
	return val
}
