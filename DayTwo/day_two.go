package main

import (
	"fmt"
	"Golang/WeekOne/Math"
)

func minicalc(x, y int ) {
	fmt.Println("Sum:",math .Add(x,y))
	fmt.Println("Difference:", math .Substract(x,y))
	fmt.Println("Product:", math .Multiply(x,y))
	result,err:=math .Divide(x,y)
	if(err!=nil){
		fmt.Println("Error:",err)
	}
	fmt.Println("Quotient:",result)
}

func main() {
	var num1, num2 int
	fmt.Println("enter first number")
	fmt.Scanln(&num1)
	fmt.Println("enter Second number")
	fmt.Scanln(&num2)
	minicalc(num1,num2)
}



