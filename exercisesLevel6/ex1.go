package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}


func foo()int{
	return 18
}
func bar()(int,string){
	return 1337, "Haxor"
}