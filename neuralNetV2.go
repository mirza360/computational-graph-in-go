package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

/*
type node struct {
	variable float64
	weight   float64
	bias     float64
	next     *node
}
*/
func main() {

	var a float64
	fmt.Print("Enter your hyper parameter a: ")
	fmt.Scan(&a)

	var b float64
	fmt.Print("Enter your hyper parameter b: ")
	fmt.Scan(&b)

	var learnigRate float64
	fmt.Print("Enter your learning Rate: ")
	fmt.Scan(&learnigRate)

	// Open csv file using Open function importing os module
	csvFile, err := os.Open("data.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	// Create CSV reader
	csvReader := csv.NewReader(csvFile)
	// Specify number of columns per row
	csvReader.FieldsPerRecord = 2
	// Specify delimiter
	csvReader.Comma = ','

	inputs, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var arrSize int
	arrSize = len(inputs) - 1

	data := make([][2]float64, arrSize)
	//fmt.Println(data)

	var i, j int
	var f string

	for i = 1; i < 4; i++ {
		for j = 0; j < 2; j++ {
			f = inputs[i][j]
			s, err := strconv.ParseFloat(f, 64)
			if err != nil {
				log.Fatal(err)
			}
			//fmt.Print(s, ", ")
			data[i-1][j] = s
		}
	}
	var row int = 3
	var gradescAData [3]float64
	gradescAData[1] = 999999999999
	var gradescA float64 //= 50000000000
	//var gradescB float64

	//check = 1 for down Check = 2 for up
	var check int = 2
	var run int = 1
	for check != 0 {
		var x float64
		var y float64

		var gradA float64 = 0
		var gradB float64 = 0
		var i int
		for i = 0; i < row; i++ {
			x = data[i][0]
			y = data[i][1]
			var lossA, lossB float64 = train(a, b, x, y)
			gradA = gradA + lossA
			gradB = gradB + lossB

		}
		gradA = math.Round(gradA*100) / 100
		//println(gradA)
		if check == 2 {
			gradescAData[2] = gradA
			if gradescAData[1] >= gradescAData[2] {
				a = a + learnigRate
				gradescAData[1] = gradescAData[2]
			} else if gradescAData[1] <= gradescAData[2] {
				a = a - (2 * learnigRate)
				check = 1
			}
		} else if check == 1 {
			gradescAData[0] = gradA
			if gradescAData[1] >= gradescAData[0] {
				a = a - learnigRate
				gradescAData[1] = gradescAData[0]
				gradescAData[1] = gradescAData[0]
			} else if gradescAData[1] <= gradescAData[0] {
				a = a + (2 * learnigRate)
				check = 2
			}
		}
		if gradescAData[0] > gradescAData[1] && gradescAData[2] > gradescAData[1] {
			check = 0
			gradescA = gradescAData[1]
		}

		print("loss ")
		print(gradescAData[1])
		print("\nWeight ")
		print(a)

		run++

	}
	println("\n")
	print("Final Loss")
	println(gradescA)
	print("\nFinal Weight")
	print(a)
	//println(gradB)

	//var MSE float64

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
	costA = math.Abs(costA) + math.Abs(a) + math.Abs(b)
	costB = math.Abs(costB) + math.Abs(a) + math.Abs(b)
	return costA, costB

}
