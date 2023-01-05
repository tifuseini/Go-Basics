package main

import "fmt"

type Circle struct {
	x, y, r float64
}

func main() {
	/** Initializating a Strucy */

	var a Circle

	b := new(Circle)

	c := Circle{x: 0, y: 0, r: 5}

	d := Circle{0, 0, 5}

	//Pointer to the struct
	e := &Circle{0, 0, 5}

	//Fields
	fmt.Println(a.x, a.y, a.r)
	fmt.Println(b.x, b.y, b.r)
	fmt.Println(c.x, c.y, c.r)
	fmt.Println(d.x, d.y, d.r)
	fmt.Println(e.x, e.y, e.r)

}
