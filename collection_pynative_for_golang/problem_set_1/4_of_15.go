// Write a program to remove characters from a string starting from zero up to n and return a new string.
package main

import "fmt"

func main() {
	word, num := getWordAndNum()
	fmt.Println(word[num:])
}

func getWordAndNum() (string, int) {
	fmt.Println("Enter a word and a number of chars to remove from the beginning of the word: ")
	var word string
	var num int
	_, err := fmt.Scan(&word, &num)
	handleErr(err)
	return word, num
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
