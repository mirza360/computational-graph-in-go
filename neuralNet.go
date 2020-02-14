package main

/*
type node struct {
	variable float64
	weight   float64
	bias     float64
	next     *node
}
*/
func main() {

	var a float64 = 1
	var b float64 = 1
	var x float64 = 3
	var y float64 = 5

	//var MSE float64

	var lossA, lossB float64 = train(a, b, x, y)
	println("Here is the loss")
	println(lossA)
	println(lossB)

}
func train(a, b, x, y float64) (float64, float64) {
	//forward pass initiation
	var i [6]float64
	i[1] = a * x
	i[2] = i[1] + b
	i[3] = i[2]
	i[4] = y - i[3]
	i[5] = i[4] * i[4]

	//backpropagation
	//differential stuff
	var diff [5]float64
	//di5/di4
	diff[4] = 2 * i[4]
	//di4/di3
	diff[3] = -1
	//di3/di2
	diff[2] = 1
	//di2/di1 	useful only for di5/da
	diff[1] = b
	//overlap cost
	var c int
	var overlapCost float64 = 1
	for c = 2; c < 5; c++ {
		overlapCost = overlapCost * diff[c]
	}
	//di1/db=x diff[1] has di2/di1
	var costA float64 = overlapCost * diff[1] * x

	//di1/db=i1
	var costB float64 = overlapCost * i[1]

	//end of differentiation stuff

	//loss function
	//var loss float64 = i[5] + math.Abs(a) + math.Abs(b)
	return costA, costB

}
