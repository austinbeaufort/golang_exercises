// Remove empty strings from a list
package main

import "fmt"

func Solve14() {
	names := [7]string{"Emma", "John", "", "Kelly", "", "Eric", ""}
	fmt.Println(removeEmptyStings(names))
}

func removeEmptyStings(names [7]string) []string {
	filteredVals := []string{}
	for _, name := range names {
		if name != "" {
			filteredVals = append(filteredVals, name)
		}
	}
	return filteredVals
}
