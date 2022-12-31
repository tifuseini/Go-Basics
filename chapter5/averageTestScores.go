package main

import "fmt"

func main() {

	//var testScores [5]float64
	//testScores[0] = 98
	//testScores[1] = 98
	//testScores[2] = 90
	//testScores[3] = 78
	//testScores[4] = 9

	testScores := [5]float64{12, 98, 89, 90, 76}

	var totalTestScores float64 = 0
	for index := 0; index < 5; index++ {
		totalTestScores += testScores[index]

	}

	average := totalTestScores / float64(len(testScores))

	fmt.Println(average)

}
