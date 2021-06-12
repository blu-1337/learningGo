package main

import "fmt"

func main() {

	x := [][]string{{"James", "Bond", "Shaken, not stirred"},{"Miss", "Moneypenny", "Hellooooo, James"}}
	fmt.Println(x)

	for i,v := range x{
		fmt.Println(i,v)
	}
	fmt.Println("_____________")

	for i,v := range x{
		for j := range v{
			fmt.Println(i,v,v[j])
		}
	}
}
