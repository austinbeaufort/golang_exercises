// Print a date in the following format: Tuesday 25 February 2020
package main

import (
	"fmt"
	"strings"
	"time"
)

func Solve4() {
	givenDate := time.Date(2020, 2, 25, 0, 0, 0, 0, time.Local)
	dateStr := givenDate.Format(time.RFC1123Z)
	fmt.Println(formatDateStr(dateStr))
}

func formatDateStr(date string) string {
	arr := strings.Split(date, " ")
	day := getFullDay(arr[0])
	numDay := arr[1]
	month := getFullMonth(arr[2])
	year := arr[3]
	return fmt.Sprintf("%s %s %s %s", day, numDay, month, year)
}

func getFullMonth(month string) string {
	for key, value := range GetMonths() {
		if month == key {
			return value
		}
	}
	panic("Month does not exist")
}

func getFullDay(val string) string {
	day := strings.TrimSuffix(val, ",")
	for key, value := range GetDaysOfWeek() {
		if day == key {
			return value
		}
	}
	panic("Day does not exist")
}
