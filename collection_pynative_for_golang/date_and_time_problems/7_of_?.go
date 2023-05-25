// print the current time in milliseconds
package main

import (
	"fmt"
	"time"
)

func Solve7() {
	fmt.Println(time.Now().UnixMilli())
}
