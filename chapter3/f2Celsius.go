package main

import "fmt"

func main() {

	fmt.Println("Enter Fahrenheit to be converted")
	var fahrenheit float64
	fmt.Scanf("%f", &fahrenheit)

	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Println("From celsuis to fahrenheit", celsius)

}
