// Print the following pattern
// 1
// 2 2
// 3 3 3
// 4 4 4 4
// 5 5 5 5 5
package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println(getLine(i))
	}
}

func getLine(num int) string {
	line := ""
	for i := 1; i <= num; i++ {
		line += fmt.Sprintf("%d ", num)
	}
	return line
}
