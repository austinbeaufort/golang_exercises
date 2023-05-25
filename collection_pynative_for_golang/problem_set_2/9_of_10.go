// Write a program to check if the given file is empty or not
package main

import (
	"fmt"
	"os"
)

func Solve9() {
	data, err := os.ReadFile("file_empty.txt")
	HandleErr(err)
	isEmpty := len(string(data)) == 0

	if isEmpty {
		fmt.Println("Empty File")
	} else {
		fmt.Println("File contains Characters")
	}
}
