package main

import (
	"fmt"
	"math"
)

func noDuplicates(nums []int )[]int{
	exist:=make(map[int]bool)
	result:=[]int{}

	for _,n:= range nums{
		if !exist[n]{
			exist[n]=true
			result = append(result, n)
		}
	}
	return result
}

func MaxMin(nums[]int)(max,min int){

	max=math.MinInt64
	min=math.MaxInt64

	for _,n:= range nums{
		if n>=max{
			max=n
		}
		 if n<=min{
			min=n
		}
	}
	return max,min
}

func main() {

	students:=map[string]int{
		"ben":78,
		"luke":99,
		"sofia":98,
		"bum":48,
	}
	scores:=[]int{}
	students["kevin"]=100
	for k,v:= range students{

		fmt.Println(k," scored",v,"marks")
		scores=append(scores,v)
	}
	max,min :=MaxMin(scores)
	fmt.Printf("Max mark: %d min mark: %d \n",max,min)
	


	numbers := []int{10,101,10}
	a,b :=MaxMin(numbers)
	fmt.Printf("Max : %d min : %d \n",a,b)
	fmt.Println(noDuplicates(numbers))
}