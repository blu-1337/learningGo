package main

import "fmt"

func main() {
	x := []int{1,2,3,4,5}

	fmt.Println(foo(x...))
	fmt.Println(bar(x))
}


func foo(x ...int)int{
	var sum int
	for _,i:=range x{
		sum += i
	}
	return sum
}
func bar(x[]int)(int){
	var sum int
	for _,i:=range x{
		sum += i
	}
	return sum
}