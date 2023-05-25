// Write a program to count occurrences of all characters within a string
// don't count spaces or newlines
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Solve10() {
	fmt.Println("Enter a String: ")
	inputReader := bufio.NewReader(os.Stdin)
	val := getNo10Input(inputReader)
	fmt.Println("\n\n================ Chars and Counts ================")
	printMapCounts(getCharCounts((val)))
}

func printMapCounts(counts map[string]int) {
	for key, value := range counts {
		fmt.Printf("Char: %s Count: %d\n", key, value)
	}
}

func getCharCounts(val string) map[string]int {
	uniques := getUniques(val)
	counts := map[string]int{}
	for _, char := range uniques {
		value := strings.Count(val, string(char))
		counts[string(char)] = value
	}
	return counts
}

func getUniques(val string) (uniques string) {
	for _, char := range val {
		if !strings.Contains(uniques, string(char)) && char != '\n' && char != ' ' {
			uniques += string(char)
		}
	}
	return
}

func getNo10Input(inputReader *bufio.Reader) string {
	val, err := inputReader.ReadString('\n')
	Check(err)
	return val
}
