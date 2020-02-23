package main

import (
	"fmt"
	"math"
)

var a float64 = 0.0
var b float64 = 0.0
var c float64 = 0.0
var d float64 = 0.0
var x float64 = 0.0
var learnigRate float64 = 0.0

type node struct {
	idx       int
	varA      float64
	varB      float64
	varTemp   float64
	resVal    float64
	diffVal   float64
	str       string
	prev      *node
	next      *node
	childDown *node
	childUp   *node
}

type list struct {
	tail *node
	head *node
}

func main() {
	//solving for y=(c(ax+b)^2+d)
	/*
		var a float64 = 1.5
		var b float64 = 1.7
		var c float64 = 1
		var d float64 = 1
		var learnigRate float64 = 0.5
		link := node{}
		head := root{}
	*/
	//head.head:=link
	//head.tail:=link

	var data [3][2]float64
	data[0][0] = float64(2)
	data[0][1] = float64(17)
	data[1][0] = float64(3)
	data[1][1] = float64(37)
	data[2][0] = float64(4)
	data[2][1] = float64(65)
	//println(gradB)

	//var MSE float64

	var a float64
	var x float64
	a = 2
	x = 3
	//var m float64 = 0
	//var di float64 = 0
	m, di, st := sqr(x, a, x, "x")
	println(m)
	println(di)
	println(st)

	item := calculate()

	item.Display()

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

func (l *list) addBottomNode(newNode *node) {
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		currentNode := l.head
		for currentNode.childDown != nil {
			currentNode = currentNode.childDown
		}
		newNode.childUp = currentNode
		currentNode.childDown = newNode
		l.tail = newNode
	}
}

func (l *list) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("idx = %v and resVal = %v and str = %#v and prev = %v and next= %v\n", list.idx, list.resVal, list.str, list.prev, list.next)
		list = list.next
	}
	fmt.Println()
}

func calculate() *list {
	//solving for y=(c(ax+b)^2+d)
	items := &list{}

	i1, df, s := mult(a, x, x, "x")
	r1, df1, s1 := mult(a, x, a, "a")
	node1 := node{idx: 1, resVal: i1, varTemp: r1}
	node1bottomNode1 := node{diffVal: df, str: s}
	node1bottomNode2 := node{diffVal: df1, str: s1}
	items.addNode(&node1)
	items.addBottomNode(&node1bottomNode1)
	items.addBottomNode(&node1bottomNode2)

	i2, df2, s2 := add(i1, b, b, "b")
	r2, df3, s3 := add(i1, b, i1, "i1")
	node2 := node{idx: 2, resVal: i2, varTemp: r2}
	node2bottomNode1 := node{diffVal: df2, str: s2}
	node2bottomNode2 := node{diffVal: df3, str: s3}
	items.addNode(&node2)
	items.addBottomNode(&node2bottomNode1)
	items.addBottomNode(&node2bottomNode2)

	i3, df4, s4 := sqr(i2, 2, i2, "i2")
	node3 := node{idx: 3, resVal: i3, diffVal: df4, str: s4}
	items.addNode(&node3)

	i4, df5, s5 := mult(c, i3, i3, "i3")
	r3, df6, s6 := mult(c, i3, c, "c")
	node4 := node{idx: 4, resVal: i4, varTemp: r3}
	node4bottomNode1 := node{diffVal: df5, str: s5}
	node4bottomNode2 := node{diffVal: df6, str: s6}
	items.addNode(&node4)
	items.addBottomNode(&node4bottomNode1)
	items.addBottomNode(&node4bottomNode2)

	return items
}

func mult(a, b, d float64, str string) (float64, float64, string) {
	var multAB float64 = a * b
	//diff
	var diffAB float64
	if a == b {
		diffAB = b
	} else if b == d {
		diffAB = a
	}
	str1 := str
	return multAB, diffAB, str1

}
func add(a, b, d float64, str string) (float64, float64, string) {
	var addAB float64 = a + b
	//diff
	var diffAB float64
	if a == b {
		diffAB = a / math.Abs(a)
	} else if b == d {
		diffAB = b / math.Abs(b)
	}
	str1 := str
	return addAB, diffAB, str1

}
func sqr(a, b, d float64, str string) (float64, float64, string) {
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

	str1 := str

	return sqrAB, diffAB, str1
}
