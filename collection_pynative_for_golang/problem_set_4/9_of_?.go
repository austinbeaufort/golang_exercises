// find the minimum and maximum values of a list
package main

import "fmt"

func Solve9() {
	list := [7]int{87, 45, 41, 65, 94, 41, 99}
	minAndMax := minMax(list)
	fmt.Printf("Min: %d\nMax: %d\n", minAndMax["min"], minAndMax["max"])
}

func minMax(nums [7]int) map[string]int {
	min, max := 0, 0

	for i, num := range nums {
		min, max = setInitialVals(map[string]int{"min": min, "max": max, "num": num, "index": i})
		min = setMin(min, num)
		max = setMax(max, num)
	}

	return map[string]int{
		"min": min,
		"max": max,
	}
}

func setInitialVals(vals map[string]int) (int, int) {
	if vals["index"] == 0 {
		return vals["num"], vals["num"]
	}
	return vals["min"], vals["max"]
}

func setMax(max, num int) int {
	if num > max {
		return num
	}
	return max
}

func setMin(min, num int) int {
	if num < min {
		return num
	}
	return min
}
