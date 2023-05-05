// merge two maps into one
package main

import "fmt"

type Nums map[string]int

func Solve13() {
	nums1 := Nums{
		"Ten":    10,
		"Twenty": 20,
		"Thirty": 30,
	}
	nums2 := Nums{
		"Fourty": 40,
		"Fifty":  50,
		"Sixty":  60,
	}

	numsList := [2]Nums{nums1, nums2}
	fmt.Println(mergeMaps(numsList))
}

func mergeMaps(numsList [2]Nums) Nums {
	newMap := make(Nums)
	for _, group := range numsList {
		for key, value := range group {
			newMap[key] = value
		}
	}
	return newMap
}
