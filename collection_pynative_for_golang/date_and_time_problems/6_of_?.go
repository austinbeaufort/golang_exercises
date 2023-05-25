// Add a week (7 days) and 12 hours to a given date
package main

import (
	"fmt"
	"time"
)

func Solve6() {
	givenDate := time.Date(2020, 3, 22, 10, 0, 0, 0, time.Local)
	newDate := givenDate.AddDate(0, 0, 7)
	newDate = newDate.Add(time.Hour * 12)
	fmt.Println(newDate)
}
