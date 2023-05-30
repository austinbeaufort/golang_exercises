// Repeat a given string str (first argument) for num times (second argument).
// Return an empty string if num is not a positive number.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Solve5() {
	fmt.Println("Enter a string: ")
	val := strings.TrimSuffix(getStringInput(), "\n") 
	fmt.Println("Enter an int: ")
	num := getInt()
	fmt.Println(repeat(val, num))
}

func repeat(val string, num int) (str string) {
	if num <= 0 {
		return
	}
	for i := 0; i < num; i++ {
		str += val
	}
	return
}

func getInt() (num int) {
	_, err := fmt.Scan(&num)
	check(err)
	return num
}

func getStringInput() string {
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	check(err)
	return input
}