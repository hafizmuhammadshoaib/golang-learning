package main

import "fmt"

func main(){
	rect := Rectangle{10,10,10,10}
	fmt.Println(rect.width)
}
type Rectangle struct {
	leftX float64
	topY float64
	height float64
	width float64
}
func (rect *Rectangle) area() float64  {
	return rect.width * rect.height
}