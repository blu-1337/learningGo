package main

import "fmt"

type blurex int
var x blurex

func main() {

	fmt.Println(x)
	fmt.Printf("%T \n", x)
	x = 42
	fmt.Println(x)

}