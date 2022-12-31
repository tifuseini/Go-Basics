package main

import "fmt"

func main() {

	var testScores [5]float64
	testScores[0] = 98
	testScores[1] = 98
	testScores[2] = 90
	testScores[3] = 78
	testScores[4] = 9

	var totalTestScores float64 = 0
	for index := 0; index < 5; index++ {
		totalTestScores += testScores[index]

	}

	average := totalTestScores / 5

	fmt.Println(average)

}
