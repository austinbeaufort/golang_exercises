// subtract seven days from a given date
package main

import (
	"fmt"
	"time"
)

func Solve3() {
	givenDateStr := "2020-02-24"
	layout := "2006-01-02"
	givenDate, err := time.Parse(layout, givenDateStr)
	HandleErr(err)
	newTime := givenDate.AddDate(0, 0, -7)
	fmt.Println(newTime.Format(layout))
}
