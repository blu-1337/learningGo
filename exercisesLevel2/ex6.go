package main

import "fmt"

func main() {
	const (
		this = 2021 + iota
		next
		afternext
		nextafternext
	)
	fmt.Println(this)
	fmt.Println(next)
	fmt.Println(afternext)
	fmt.Println(nextafternext)
}
