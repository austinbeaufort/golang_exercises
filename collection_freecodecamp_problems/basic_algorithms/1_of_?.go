// Reverse the provided string and return the reversed string.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Solve1() {
	fmt.Println("Enter a string or sentence: ")
	inputReader := bufio.NewReader(os.Stdin)
	result := reverse(strings.Trim(get_input(inputReader), "\n"))
	fmt.Println(result)

	fccChecks()
}

func reverse(val string) string {
	newString := ""
	for i := len(val) - 1; i >= 0; i-- {
		newString += string(val[i])
	}
	return newString
}

func get_input(inputReader *bufio.Reader) string {
	input, err := inputReader.ReadString('\n')
	check(err)
	return input
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func fccChecks() {
	// items that must pass in freecodecamp
	res1 := reverse("hello") == "olleh"
	res2 := reverse("Howdy") == "ydwoH"
	res3 := reverse("Greetings from Earth") == "htraE morf sgniteerG"
	fmt.Println("========== FCC CHECKS ==========")
	if res1 && res2 && res3 {
		fmt.Println("success - fcc checks passed!")
	} else {
		panic("failed - fcc checks")
	}
}
