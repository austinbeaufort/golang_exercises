// Write a program to take three names as input from a user in the single Scan() function call.
package main

import "fmt"

func Solve7() {
	fmt.Println("Enter three names: ")
	var name1, name2, name3 string
	_, err := fmt.Scan(&name1, &name2, &name3)
	HandleErr(err)

	fmt.Printf("\nNames Entered:\n")
	fmt.Printf("Name1: %s\n", name1)
	fmt.Printf("Name2: %s\n", name2)
	fmt.Printf("Name3: %s\n", name3)
}
