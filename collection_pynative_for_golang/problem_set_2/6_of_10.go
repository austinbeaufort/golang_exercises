// Write all content of a given file into a new file by skipping line number 5
// Uses file.txt, creates new_file.txt
package main

import (
	"fmt"
	"os"
	"strings"
)

func Solve6() {
	data, err := os.ReadFile("file.txt")
	HandleErr(err)
	formattedData := formatFileData(data)
	err = os.WriteFile("new_file.txt", []byte(formattedData), 0644)
	HandleErr(err)
	fmt.Println("new file written")
}

func formatFileData(data []byte) string {
	dataArr := strings.Split(string(data), "\n")
	dataSlice := []string{}
	dataSlice = append(dataSlice, dataArr[:4]...)
	dataSlice = append(dataSlice, dataArr[5:]...)
	return strings.Join(dataSlice, "\n")
}
