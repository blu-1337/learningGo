package main

import "fmt"

func main() {
	x :=  []int{1,2,4,5,23,55,33,1992,3491}

	for i:=0; i<len(x); i++{
		fmt.Println(x[i])
	}

	y := []int{9,99,999,9999}
	x = append(x,y...)
	fmt.Println(x)

	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
