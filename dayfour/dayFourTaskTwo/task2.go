package main

import (
	"fmt"
	"math"
)

//circle structure

type circle struct {
	Radius float64
}


func (c circle) Area() float64 {
	return math.Pi * c.Radius* c.Radius
}
func (c circle) perimeter() float64{
	return 2* math.Pi * c.Radius
}

//rectangel structure

type rect struct{
	Length 	float64
	Width 	float64
}


func (r rect) Area() float64 {
	return  r.Length* r.Width
}
func (r rect) perimeter() float64{
	return 2 * (r.Length + r.Width)
}

//shape interface

type shape interface{
	Area() 			float64
	perimeter()		float64
}

func PrintDetails(s shape){
	switch v:= s.(type){
	case circle:
		fmt.Printf("Area of the circle with radius, %.2f is: %.2f, and its perimiter is: %.2f \n",v.Radius,v.Area(),v.perimeter())
	case rect:
		fmt.Printf("Area of the rectangle is: %.2f, and its perimiter is: %.2f \n",v.Area(),v.perimeter())
	default:
		fmt.Println("unknown shape ")
	}

}

func main() {

	c1:=circle{2.34}
	r1:=rect{3.45,4.35}
	PrintDetails(c1)
	PrintDetails(r1)
}
