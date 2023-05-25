// Read fourth line of a given file
package main

import (
	"bufio"
	"fmt"
	"os"
)

func Solve10() {
	file, err := os.Open("file.txt")
	HandleErr(err)
	defer file.Close()

	lineNum := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if lineNum == 4 {
			fmt.Printf("Line 4: %s\n", scanner.Text())
		}
		lineNum += 1
	}
	HandleErr(scanner.Err())
}
