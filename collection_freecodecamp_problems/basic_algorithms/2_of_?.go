// Return the length of the longest word in the provided sentence.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Solve2() {
	fmt.Println("Enter a sentence to find the longest word: ")
	input := getInput()
	longestWord := getLongestWord(input)
	fmt.Printf("\nThe longest word is: %s\n", longestWord)
	fmt.Printf("The length of the word is: %d\n\n", len(longestWord))

	fccChecks1()
}

func getLongestWord(sentence string) string {
	words := strings.Split(sentence, " ")

	longest := ""
	for _, word := range words {
		if len(word) > len(longest) {
			longest = word
		}
	}
	trimmed := strings.Trim(longest, "\n")
	return trimmed
}

func getInput() string {
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	check(err)
	return input
}


func fccChecks1() {
	// items that must pass in freecodecamp
	res1 := len(getLongestWord("The quick brown fox jumped over the lazy dog")) == 6
	res2 := len(getLongestWord("May the force be with you")) == 5
	res3 := len(getLongestWord("Google do a barrel roll")) == 6
	res4 := len(getLongestWord("What is the average airspeed velocity of an unladen swallow")) == 8
	res5 := len(getLongestWord("What if we try a super-long word such as otorhinolaryngology")) == 19
	if res1 && res2 && res3 && res4 && res5 {
		fmt.Println("success - fcc checks passed!")
	} else {
		panic("failed - fcc checks")
	}
}