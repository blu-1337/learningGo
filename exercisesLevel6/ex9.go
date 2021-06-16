package main

import "fmt"

func sum(x ...int) int{
	var s int
	for _, elem := range x{
		s += elem
	}
	return s
}

func unevenSum(hey func(x ...int) int, y ...int) int{ //functia f pentru suma, elementele testate

	var newList []int
	for _,elem := range y{
		if elem % 2 != 0{
			newList = append(newList, elem)
		}
	}
	return hey(newList...)

}

func main() {
	x :=[]int{1,2,3,4,5}
	fmt.Println(sum(x...)) //unfurling
	fmt.Println(unevenSum(sum,x...))
}
