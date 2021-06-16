package main

import "fmt"

func main() {
	p1:= struct {
		firstName string
		occupation string
		age int
	}{
		firstName: "Cristian",
		occupation: "Programmer",
		age: 27,
	}
	fmt.Println(p1)
	fmt.Println(p1.age)
}
