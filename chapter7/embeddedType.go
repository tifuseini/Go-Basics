package main

import "fmt"

type Person struct {
	name string
	age  string
}

func (p *Person) Talk() {
	fmt.Println("Hi my name is ", p.name)
	fmt.Println("I am  ", p.age)
}

type Android struct {
	Person Person
	Model  string
}

func main() {

	a := new(Android)
	a.Person.Talk()

}
