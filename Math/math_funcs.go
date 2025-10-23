package math 

import("fmt")

func Add(a, b int) int {
	return a + b
}

func Divide(a, b int) (int,error) {
	if b == 0 {
		 return 0, fmt.Errorf("%d cannot be divided by %d",a,b)
	}
	return a/b, nil
}

func Multiply(a,b int) int{
	return a*b
}
func Substract(a,b int)int{
	return a-b
}
