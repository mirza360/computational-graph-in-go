package main

import (
	"fmt"
)

type node struct {
	val int
}

func main() {
	fmt.Println("test")
	var arrRow, arrCol, i, j int
	fmt.Println("Enter your row size: ")
	fmt.Scan(&arrRow)
	fmt.Println("Enter your column size: ")
	fmt.Scan(&arrCol)
	table := sliceTable(arrRow, arrCol)

	for i = 0; i < arrRow; i++ {
		for j = 0; j < arrCol; j++ {
			println(table[i][j].val)
		}
	}
}

func sliceTable(arrRow, arrCol int) [][]node {
	var i, j int
	data := make([][]node, arrRow)

	println(len(data))
	count := 1
	for i = 0; i < arrRow; i++ {
		data[i] = make([]node, arrCol)
		for j = 0; j < arrCol; j++ {
			data[i][j] = node{val: count}
			count++
		}
	}
	return data
}
