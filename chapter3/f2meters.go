package main

import "fmt"

func main() {

	fmt.Println("Enter feet to be converted")
	var feet float64;
	fmt.Scanf("%f", &feet)

	meter := feet * 0.3048
	fmt.Println("Update feet", meter)

}
