// Find the day of the week of a given date
package main

import (
	"fmt"
	"strings"
	"time"
)

func Solve5() {
	givenDate := time.Date(2020, 7, 26, 0, 0, 0, 0, time.Local)
	formattedDate := givenDate.Format(time.RFC850)
	arr := strings.Split(formattedDate, " ")
	fmt.Println(strings.TrimSuffix(arr[0], ","))
}
