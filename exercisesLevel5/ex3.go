package main

import "fmt"

type vehicle struct{
	doors int
	color string
}
type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	car1 := sedan{
		vehicle: vehicle{doors: 2,
			color: "blue",},
		luxury: true,
	}
	car2 := truck{
		vehicle: vehicle{doors: 5,
			color: "orange",},
		fourWheel: false,
	}

	fmt.Println(car1)
	fmt.Println(car2)
	fmt.Println(car1.color)
	fmt.Println(car2.fourWheel)
}
