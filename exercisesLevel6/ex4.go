package main

import "fmt"

type person struct{
	first string
	last string
	age int
}
func (p person) speak(){
	fmt.Println("My name is", p.last, p.first, "and I am ", p.age, "years old!")
}


func main() {
	p1 := person{
		first: "Cristian",
		last:"Gîlcă",
		age: 1337,
	}
	p1.speak()
}
