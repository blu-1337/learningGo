package main

import "fmt"

func main() {
	x:=[]int{42,43,44,45,46,47,48,49,50,51}

	fmt.Println(x)
	y:=x[:5]
	fmt.Println(y)
	y = x[5:]
	fmt.Println(y)
	y = x[2:7]
	fmt.Println(y)
	y = x[1:6]
	fmt.Println(y)


}
