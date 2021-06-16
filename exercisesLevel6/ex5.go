package main

import (
	"fmt"
	"math"
)

type circle struct{
	radius float32
}

type rectangle struct{
	length float32
	width float32
}
type SHAPE interface{
	area() float32
}
func INFO(s SHAPE) float32{
	return s.area()
}

func (c circle) area() float32{
	return math.Pi * c.radius * c.radius
}
func (s rectangle) area() float32 {
	return s.length * s.width
}

func main() {
	c1 := circle {radius: 4.23}
	s1 := rectangle{length: 10, width: 20}
	fmt.Println(c1.area(),s1.area())
	fmt.Println(INFO(c1),INFO(s1))
}