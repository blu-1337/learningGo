package main

import "fmt"

func main() {

	favSport := "football"

	switch favSport {
	case "volley" , "badminton", "tennis": fmt.Println("girl sport")
	case "soccer", "football": fmt.Println("man sport")
	}
}
