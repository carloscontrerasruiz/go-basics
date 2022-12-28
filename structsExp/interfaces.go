package main

import "fmt"

type PrintInfo interface {
	getMessage() string
}

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
	endDate string
}

type TemporaryEmployee struct {
	Person
	EmployeeC
	taxRate int
}

func (t TemporaryEmployee) getMessage() string {
	return "Temporary Employee"
}

func (t FullTimeEmployee) getMessage() string {
	return fmt.Sprintf("Full time Employee %v", t.id)
}

func getMessageInterfaces(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {

	ftEmp2 := FullTimeEmployee{
		endDate: "",
		EmployeeC: EmployeeC{
			id: 108,
		},
		Person: Person{
			name: "Pancho",
		},
	}

	fmt.Println(ftEmp2.id)
	getMessageInterfaces(ftEmp2)

	teEmp := TemporaryEmployee{}
	getMessageInterfaces(teEmp)
}
