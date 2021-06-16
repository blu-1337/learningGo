package main

import "fmt"

func main()  {
	defer foo()
	bar()
}

func foo(){
	fmt.Println("I should run first!")
}

func bar(){
	fmt.Println("I should be second!")
}
