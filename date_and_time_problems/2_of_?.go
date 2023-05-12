// convert string to datetime
package main

import (
	"fmt"
	"time"
)

func Solve2() {
	layout := "January 02, 2006 3:04 PM"
	dateString := "February 25, 2020 4:20 PM"
	parsedTime, err := time.Parse(layout, dateString)
	HandleErr(err)
	fmt.Println(parsedTime)
}
