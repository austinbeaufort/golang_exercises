// Write a program to accept a string from the user and display characters that are present at an even index number.
package main

import "fmt"

func main() {
	word := getWord()
	printEvens(word)
}

func printEvens(word string) {
	for index, char := range word {
		if index%2 == 0 {
			fmt.Println(string(char))
		}
	}
}

func getWord() string {
	fmt.Print("Enter a word to print chars on even indexes: ")
	var word string
	_, err := fmt.Scan(&word)
	handleErr(err)
	return word
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
