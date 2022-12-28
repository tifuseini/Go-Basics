package main

import "fmt"

func main() {
	/** AND */
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && false)

	/** OR */
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || false)

	/* NOT */
	fmt.Println(!true)

	fmt.Println(!false)

}
