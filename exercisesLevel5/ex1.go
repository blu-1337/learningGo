package main

import "fmt"

type person struct{
	firstName string
	lastName string
	favoriteIceCreamFlavours []string
}

func main() {
	pers1 := person{
		firstName: "Crsitian",
		lastName: "Gîlcă",
		favoriteIceCreamFlavours: []string{"Hazelnut", "vanilla"},
	}
	pers2 := person{
		firstName: "Luisa",
		lastName: "Betancourt",
		favoriteIceCreamFlavours: []string{`mint`, `chocolate chip`},
	}
	fmt.Println(pers1, pers2)

	for i,v := range pers1.favoriteIceCreamFlavours{
		fmt.Println(i,v)
	}
	for i,v := range pers2.favoriteIceCreamFlavours{
		fmt.Println(i,v)
	}

}
