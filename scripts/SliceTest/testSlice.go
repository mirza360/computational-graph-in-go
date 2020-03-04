package main

import (
	"fmt"
)

type node struct {
	val int
}

func main() {
	fmt.Println("test")

	var arrRow, arrCol int
	fmt.Println("Enter your row size: ")
	fmt.Scan(&arrRow)
	fmt.Println("Enter your column size: ")
	fmt.Scan(&arrCol)
	data := sliceTable(arrRow, arrCol)

	fmt.Println(data[3][3])
}

func sliceTable(arrRow, arrCol int) [][]node {
	println(arrRow, " : ", arrCol)
	data := make([][]node, arrRow, arrCol)
	return data
}
