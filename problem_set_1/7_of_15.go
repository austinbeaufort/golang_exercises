// Return the count of a given substring from a string
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	sentenceMsg := "Enter a sentence: "
	subStringMsg := "Enter a substring to count no of occurences within the given setence: "
	sentence := getInput(inputReader, sentenceMsg)
	subString := removeNewline(getInput(inputReader, subStringMsg))
	fmt.Printf("Your substring occurs %d times within the given sentence.\n", strings.Count(sentence, subString))
}

func removeNewline(str string) string {
	newlineIndex := len(str) - 1
	return str[:newlineIndex]
}

func getInput(inputReader *bufio.Reader, msg string) string {
	fmt.Println(msg)
	input, err := inputReader.ReadString('\n')
	handleErr(err)
	return input
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
