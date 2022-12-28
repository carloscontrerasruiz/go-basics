package main

import "fmt"

type Person struct {
	name string
	age  int
}
type EmployeeC struct {
	id int
}
type FullTimeEmployee struct {
	Person
	EmployeeC
}

func main() {
	ftEmp := FullTimeEmployee{}

	ftEmp.name = "Carlos"
	ftEmp.id = 1

	fmt.Println(ftEmp)

	ftEmp2 := FullTimeEmployee{
		Person{
			name: "Pancho",
		},
		EmployeeC{
			id: 199,
		},
	}

	fmt.Println(ftEmp2)
}
