// Print multiplication table form 1 to 10
package main

import "fmt"

func main() {
	fmt.Println("****** MULTIPLICATION TABLE ******")
	for i := 1; i <= 10; i++ {
		fmt.Println(getRow(i))
	}
}

func getRow(num int) string {
	row := ""
	for i := 1; i <= 10; i++ {
		newStr := formatNum(num, i)
		row += newStr
	}
	return row
}

func formatNum(num, i int) string {
	if num*i <= 9 {
		return fmt.Sprint(num*i) + "  "
	}
	return fmt.Sprint(num*i) + " "
}
