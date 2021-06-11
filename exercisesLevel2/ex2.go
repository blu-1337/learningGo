package main

import "fmt"

func main() {
	g := (100 == 100)
	h := (24 <= 24)
	i := ("some" >= "stri")
	j := (1 != 2)
	k := (211<12)
	l := (32>21)

	fmt.Print(g,h,i,j,k,l)
}