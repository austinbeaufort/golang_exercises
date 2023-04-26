package main

import "fmt"

func GetStrInput() (val string) {
	_, err := fmt.Scan(&val)
	Check(err)
	return
}
