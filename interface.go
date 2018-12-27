package main
import (
	. "fmt"
	_ "strings"
	"math"
)

func main()  {
	circle := Circle{4}
	rectangle := Rectangle{2, 4}
	Println("circle area", getArea(circle))
	Println("rectangle area", getArea(rectangle))
}

// create interface
type Shape interface{
	area() float64
}

type Circle struct{
	radius float64
}

type Rectangle struct{
	height float64
	width float64
}

func (c Circle) area() float64{
	return math.Pi * math.Pow(c.radius, 2)
}

func (r Rectangle) area() float64{
	return r.height * r.width
}

func getArea(s Shape) float64  {
	return s.area()
}