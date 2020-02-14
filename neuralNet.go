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

	var a float32 = 1.5
	var b float32 = 1.7
	var learnigRate float32 = 0.5

	var data [3][2]float32
	data[0][0] = float32(2)
	data[0][1] = float32(3)
	data[1][0] = float32(3)
	data[1][1] = float32(5)
	data[2][0] = float32(4)
	data[2][1] = float32(10)
	var row int = 3

	var gradescA float32 = 50000000000
	//var gradescB float32

	//check = 1 for down Check = 2 for up
	var check int = 2
	for check != 0 {
		var x float32
		var y float32

		var gradA float32 = 0
		var gradB float32 = 0
		var i int
		for i = 0; i < row; i++ {
			x = data[i][0]
			y = data[i][1]
			var lossA, lossB float32 = train(a, b, x, y)
			gradA = gradA + lossA
			gradB = gradB + lossB

		}
		if gradescA > gradA {
			gradescA = gradA
			if check == 1 {
				a = a + learnigRate
			} else {
				a = a - learnigRate
			}
		} else if gradescA < gradA {
			if check == 1 {
				check = 2
				a = a - (2 * learnigRate)
			} else {
				check = 1
				a = a + (2 * learnigRate)
			}

		} else if int(gradescA) == int(gradA) {
			check = 0

		}
		println("Here is the loss")
		println(gradescA)

	}
	println("Here is the final thing")
	println(gradescA)
	//println(gradB)

	//var MSE float32

}

func train(a, b, x, y float32) (float32, float32) {
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
	//di2/di1 	useful only for di5/da
	diff[1] = b
	//overlap cost
	var c int
	var overlapCost float32 = 1
	for c = 2; c < 5; c++ {
		overlapCost = overlapCost * diff[c]
	}
	//di1/db=x diff[1] has di2/di1
	var costA float32 = overlapCost * diff[1] * x

	//di1/db=i1
	var costB float32 = overlapCost * i[1]

	//end of differentiation stuff

	//loss function
	//var loss float32 = i[5] + math.Abs(a) + math.Abs(b)
	return costA, costB

}
