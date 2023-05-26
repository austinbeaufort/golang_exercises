// Check if a string (first argument, str) ends with the given target string (second argument, target).
// solve first with build in method, then solve without built in method.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Solve4() {
	sentence, suffix := getInputs2()

	// checking with built in method
	fmt.Println(strings.HasSuffix(sentence, suffix))

	// checking with custom method
	fmt.Println(hasSuffix(sentence, suffix))
}

func hasSuffix(sentence string, suffix string) bool {
	prefixTrimIndex := len(sentence) - len(suffix)
	sentenceSuffix := sentence[prefixTrimIndex:]
	return sentenceSuffix == suffix 
}


func getInputs2() (string, string) {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a sentence: ")
	sentence, err := inputReader.ReadString('\n') 
	check(err)
	fmt.Println("Enter an ending string to check against: ")
	suffix, err := inputReader.ReadString('\n')
	check(err)

	return sentence, suffix
}
