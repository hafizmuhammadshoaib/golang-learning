package main

import (
	"fmt"
	"math"
)
func main()  {

	rect := Rectangle{10,20}
	circ := Circle{4}
	fmt.Println(getArea(rect),getArea(circ))
}
type Shape interface{
	area() float64
}

type Rectangle struct {
	height float64
	width float64
}

type Circle struct {
	radius float64
}

func (rect Rectangle) area() float64  {
	return rect.width * rect.height
}
func (c Circle) area() float64  {
	return math.Pi * math.Pow(c.radius,2)
}
func getArea(shape Shape) float64  {
	return shape.area()
}