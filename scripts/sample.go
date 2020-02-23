package main

import "math"

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

func calculate() *node {
	//solving for y=(c(ax+b)^2+d)
	items := &list{}
	p, q, s := mult(a, x, x, "x")
	node1 := node{idx: 1, resVal: p, str: s}
	bottomNode1 := node{diffVal: q}
	items.addNode(&node1)
	items.addBottomNode(&bottomNode1)

	return items.head
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
