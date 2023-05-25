// Calculate the number of days between two given dates
package main

import (
	"fmt"
	"math"
	"time"
)

func Solve8() {
	date1 := time.Date(2020, 2, 25, 0, 0, 0, 0, time.Local)
	date2 := time.Date(2020, 9, 17, 0, 0, 0, 0, time.Local)
	diff := (date2.Sub(date1) / 24).Hours()
	fmt.Printf("%d days\n", int64(math.Round(diff)))
}
