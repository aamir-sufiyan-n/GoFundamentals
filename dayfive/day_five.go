package main

import ("fmt")

type student struct{
	Name 	string
	Age		int
}

func listStudents(s []student){

	for _,student:= range s{
		fmt.Printf("Name: %s, Age: %d \n",student.Name, student.Age)
	}
}
func AddStudent(s *[]student, n string, age int){
	newStudent:=student{n,age}
	 *s=append(*s, newStudent)
}
func updateStudent(index int, n string, a int , s *[]student){

	if(index<=0|| index>len(*s)){
	return}
	(*s)[index-1].Name=n
	(*s)[index-1].Age=a
}



func deleteStudent(index int, s *[]student){
	if(index<0||index>len(*s)){
		return
	}
	(*s)=append((*s)[:index],(*s)[index+1:]... )
}

var students []student
func main(){
    for {
        fmt.Printf("choose an action \n")
		fmt.Println("1.List students")
		fmt.Println("2.Add a student")
		fmt.Println("3.Remove a student")
		fmt.Println("4.Update a student's details")
		fmt.Println("5.EXIT")

		var choice int
		fmt.Println("enter a choice(1,2,3,4,5):")

		fmt.Scanln(&choice)
		switch choice{
		case 1:
			listStudents(students)
		case 2:
			var name string
			var age int
			fmt.Println("enter Student Details: name,age")
			fmt.Scanln(&name,&age)
			AddStudent(&students,name,age)
		case 3:
			var index int
			fmt.Println("enter student index to remove:")
			fmt.Scanln(&index)
			deleteStudent(index,&students)
		case 4:
			var name string
			var age int
			var index int
			fmt.Println("enter the student index to update")
			fmt.Scanln(&index)						
			fmt.Println("enter student updated details:name,age")
			fmt.Scanln(&name,&age)						
			updateStudent(index,name,age,&students)
		case 5:
			fmt.Println("exiting.....")
			return
		default:
			fmt.Println("invalid choice")
			}
		
    }
    
}