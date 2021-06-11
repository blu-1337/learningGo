package main

import "fmt"

func main() {
	x := 1994
	for {
		if x>= 1994 && x<= 2021 {
			fmt.Println(x)
			x++
		} else if x == 2022 {
			fmt.Println("last one cowboy")
			x++
		} else {break}
	}
}
