package main

import "fmt"

func main() {
	for i:=65; i<91; i++ {
		fmt.Println(i)
		fmt.Printf("\t%U %q\n", i,i)
		fmt.Printf("\t%U %q\n", i,i)
		fmt.Printf("\t%U %q\n", i,i)
	}
}
