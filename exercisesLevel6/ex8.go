package main

import "fmt"

func iReturnAFunc() func()int{
	return func()int{
		return 1337
	}
}

func main() {
	x:=iReturnAFunc()
	fmt.Println(x(), "whaat")
}
