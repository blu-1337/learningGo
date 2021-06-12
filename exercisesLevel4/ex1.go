package main

import "fmt"

func main() {
	x:=[5]int{1337, 2448, 3559, 4660}
	for i, v := range x{
		fmt.Println(i, v)
	}
	fmt.Printf("%T", x)
}
