package main

import (
	"fmt"
	"math"
)

var a float64 = 1.2
var b float64 = 1.4
var c float64 = 1.6
var d float64 = 1.8
var x float64 = 2
var y float64 = 2.2
var learnigRate float64 = 0.005

type node struct {
	resVal   float64
	diffVal1 float64
	diffVal2 float64
	str      string
	flg      string
	calcDiff float64
	prev     *node
	next     *node
}

type list struct {
	tail *node
	head *node
}

func main() {

	item := nodeOps()
	item.Display()
	//item.DisplayReverse()

}

func (l *list) addNode(newNode *node) {
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		currentNode := l.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		newNode.prev = currentNode
		currentNode.next = newNode
		l.tail = newNode
	}
}

func (l *list) Display() {
	list := l.head
	//listBottom := list.childDown
	for list != nil {
		fmt.Printf("resVal = %v & diffVal1 = %v & diffVal2 = %v & str = %#v & flag = %#v & calcDiff = %v\n", list.resVal, list.diffVal1, list.diffVal2, list.str, list.flg, list.calcDiff)
		list = list.next
	}
	fmt.Println()
}

func (l *list) DisplayReverse() {
	list := l.tail
	for list != nil {
		if list.flg == "" {
			fmt.Printf("resVal = %v\n", list.resVal)
			list = list.prev
		}
	}
	fmt.Println()
}

func nodeOps() *list {
	//solving for y=(c(ax+b)^2+d)^2
	items := &list{}
	i1, df1 := sqr(x, 2, x)
	node1 := node{resVal: i1, diffVal1: df1}
	items.addNode(&node1)

	i2, df2 := mult(c, i1, i1)
	i2, df3 := mult(c, i1, c)
	node2 := node{resVal: i2, diffVal1: df2, diffVal2: df3, str: "c"}
	items.addNode(&node2)

	i3, df4 := mult(b, x, x)
	i3, df5 := mult(b, x, b)
	node3 := node{resVal: i3, diffVal1: df4, diffVal2: df5, str: "b", flg: "flag"}
	items.addNode(&node3)

	i4, df6 := add(i3, a, i3)
	i4, df7 := add(i3, a, a)
	node4 := node{resVal: i4, diffVal1: df6, diffVal2: df7, str: "a", flg: "flag"}
	items.addNode(&node4)

	i5, df8 := sqr(i4, 2, 14)
	node5 := node{resVal: i5, diffVal1: df8, flg: "flag"}
	items.addNode(&node5)

	i6, df9 := add(i2, i5, i2)
	i6, df10 := add(i2, i5, i5)
	node6 := node{resVal: i6, diffVal1: df9, diffVal2: df10}
	items.addNode(&node6)

	i7, df11 := add(y, -i6, i6)
	node7 := node{resVal: i7, diffVal1: df11}
	items.addNode(&node7)

	i8, df12 := sqr(i7, 2, i7)
	node8 := node{resVal: i8, diffVal1: df12, calcDiff: 1}
	items.addNode(&node8)

	currentNode := items.tail
	for currentNode.prev != nil {
		currentNode.prev.calcDiff = currentNode.calcDiff + currentNode.diffVal1
		currentNode = currentNode.prev
	}

	return items
}

func mult(a, b, d float64) (float64, float64) {
	var multAB float64 = a * b
	//diff
	var diffAB float64
	if a == b {
		diffAB = b
	} else if b == d {
		diffAB = a
	}

	return multAB, diffAB
}

func add(a, b, d float64) (float64, float64) {
	var addAB float64 = a + b
	//diff
	var diffAB float64
	if a == b {
		diffAB = a / math.Abs(a)
	} else if b == d {
		diffAB = b / math.Abs(b)
	}

	return addAB, diffAB
}

func sqr(a, b, d float64) (float64, float64) {
	var sqrAB float64 = a
	var c int
	for c = 1; c < int(b); c++ {
		sqrAB = sqrAB * sqrAB
		//println(sqrAB)
	}
	//println("---")
	//println(sqrAB)
	//diff
	var diffAB float64 = b * (sqrAB / a)

	return sqrAB, diffAB
}
