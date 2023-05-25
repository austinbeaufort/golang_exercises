// Convert Decimal number to octal
package main

import (
	"fmt"
	"strconv"
)

func Solve3() {
	fmt.Println("Enter a number to convert to Octal: ")
	num := GetInt()
	octal := strconv.FormatInt(int64(num), 8)
	fmt.Printf("Decimal num: %d, Converted to Octal: %s\n", num, octal)
}
