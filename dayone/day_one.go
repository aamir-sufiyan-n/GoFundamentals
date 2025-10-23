package main

import "fmt"

// fibonacci series

// func main() {
// 	var n int
// 	fmt.Println("enter where you want the series to end:")
// 	fmt.Scanln(&n)
// 	var a, b int = 0, 1
// 	for i := 0; i < n; i++ {
// 		fmt.Println(a)
// 		a, b = b, a+b
// 	}
// }

// odd or even

// func main() {
// 	var num int
// 	fmt.Print("enter a number:")
// 	fmt.Scanln(&num)
// 	if num%2 == 0 {
// 		fmt.Println("even number")
// 	} else {
// 		fmt.Println("odd number")
// 	}
// }

// switch calculator

func main() {
	var n1, n2 int
	var operator string

	fmt.Println("enter the first number")
	fmt.Scanln(&n1)

	fmt.Println("enter the Second number")
	fmt.Scanln(&n2)

	fmt.Println("enter an operator(+,-,/,*)")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		fmt.Println("Result:", n1+n2)
	case "-":
		fmt.Println("Result:", n1-n2)
	case "*":
		fmt.Println("Result:", n1*n2)
	case ("/"):
		if n2 != 0 {
			fmt.Println("Result:", n1/n2)
		} else {
			fmt.Println("cannot be devided by zero")
		}
	default:
		fmt.Println("invalid operator")
	}

}
