package main

import "fmt"

func main() {
	const (
		a int = 12
		b bool = true
	)
	const (
		c = 1
		d = 3
	)
	fmt.Printf("%T\t%T\t%T\t%T \n", a,b,c,d)
}
