package main

import "math"

//global variables
var a float64 = 1.5
var b float64 = 1.7
var lossA float64 = 72
var lossB float64 = 25

var learnigRate float64 = 0.5

func main() {

	var tempA float64
	var tempB float64
	var gradescAData [3]float64
	gradescAData[1] = 999999999999

	//var gradescB float64

	var check int = 2

	for check != 0 {

		tempA = a + learnigRate*lossA
		tempB = b + learnigRate*lossB
		//start fp and bp
		//end of fp and bp
		//get lossA and LossB from BP
		cost := lossA + lossB

		cost = math.Round(cost*100) / 100

		//println(gradA)

		gradescAData[2] = cost
		if gradescAData[1] >= gradescAData[2] {
			a = tempA
			b = tempB
			gradescAData[1] = gradescAData[2]
		} else if ((learnigRate / 1) < 1) && gradescAData[2] > gradescAData[1] {
			check = 0

		} else if gradescAData[1] <= gradescAData[2] {
			learnigRate = learnigRate * (-1)
			//check = 1

		}

	}

	//End of Gradient Descent, all parameters are updated accordingly

}
