package main

import "math"

type node struct {
	i    int
	varA float64
	varB float64
	next *node
	prev *node
type root struct {
	head *node
	tail *node
}

func main() {

	//solving for y=(c(ax+b)^2+d)

	var a float64 = 1.5
	var b float64 = 1.7
	var c float64 = 1
	var d float64 = 1
	var learnigRate float64 = 0.5
	link :=node{}
	head:=root{}
	head.head:=link
	head.tail:=link
	



	var data [3][2]float64
	data[0][0] = float64(2)
	data[0][1] = float64(17)
	data[1][0] = float64(3)
	data[1][1] = float64(37)
	data[2][0] = float64(4)
	data[2][1] = float64(65)
	//println(gradB)

	//var MSE float64

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
	for c = 0; c < int(b); c++ {
		sqrAB = sqrAB * sqrAB
	}
	//diff
	var diffAB float64 = b * (sqrAB / a)

	return sqrAB, diffAB

}
