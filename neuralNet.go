package main

/*
type node struct {
	variable float32
	weight   float32
	bias     float32
	next     *node
}
*/
func main() {

	var a float32 = 1
	var b float32 = 1
	var x float32 = 2
	var y float32 = 3
	train(a, b, x, y)

}
func train(a, b, x, y float32) {
	//forward pass initiation
	var i [6]float32
	i[1] = a * x
	i[2] = i[1] + b
	i[3] = i[2]
	i[4] = y - i[3]
	i[5] = i[4] * i[4]
	//backpropagation
	//differential stuff
	var diff [5]float32
	//di5/di4
	diff[4] = 2 * i[4]
	//di4/di3
	diff[3] = -1
	//di3/di2
	diff[2] = 1
	//di2/di1
	diff[1] = b
	//overlap cost
	var c int
	var overlapCost float32 = 1
	for c = 1; c < 5; c++ {
		overlapCost = overlapCost * diff[c]
	}
	//di1/db=x
	//var costA float32 = overlapCost * x

	//di1/db=i1
	//var costB float32 = overlapCost * i[1]
	for c = 1; c < 6; c++ {
		println(i[c])

	}
	println()
	for c = 1; c < 5; c++ {
		println(diff[c])

	}

	//println(costA)
	//println(costB)

}
