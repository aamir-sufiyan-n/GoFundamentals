// person structure with a display method

package main

import "fmt"

type person struct {
	Name  string
	Age   int
	Place string
}

func (p person) display() {
	fmt.Printf("%s is %d years old and lives in %s \n",p.Name,p.Age,p.Place)
}

func NewPerson(name string ,age int, place string  ) person{

	return person{
		Name: name,
		Age: age,
		Place: place,
	}
}

	func main() {
		p1:=NewPerson("Aamir",18,"kollam")
		p2:=person{"joji",21,"thrissur"}
		p3:=person{}
		p4:=person{
			Name: "abhijith",
			Age: 21,
			Place: "kottayam",
		}
		p1.display()
		p3.display()
		p4.display()
		p2.display()
	}

//a function to modify integers using a pointer

// package main

// import "fmt"


// func modifyNum(x *int,y int){
// 	fmt.Printf("Previous: %d, Modified:%d",*x,*x*y)
// }


// func main(){
// 	num:=10
// 	modifyNum(&num,8)
// }