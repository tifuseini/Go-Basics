package main

import "fmt"

func main() {

	for index := 1; index <= 10; index++ {

		switch index {
		case 1:
			fmt.Println("one")
		default:
			fmt.Println("Unknown Number")

		}

	}

}
