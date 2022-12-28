package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func (e *Employee) setName(name string) {
	e.name = name
}

func (e *Employee) setId(id int) {
	e.id = id
}

func (e *Employee) GetId() int {
	return e.id
}

func ConstructorEmployee(is int, name string) *Employee {
	return &Employee{
		id:   is,
		name: name,
	}
}

func main() {
	emp := Employee{
		id:   5,
		name: "Carlos",
	}

	fmt.Printf("%v", emp)

	emp.name = "Pancho"
	emp.setId(109)

	fmt.Printf("%v", emp)

	//new devuleve referencia
	emp2 := new(Employee)
	fmt.Printf("%v\n", *emp2)

	emp3 := &Employee{
		id:   600,
		name: "Referencia",
	}
	fmt.Printf("%v\n", *emp3)
	fmt.Printf("%v\n", emp3)

	emp4 := ConstructorEmployee(700, "Metodo constructor")
	fmt.Printf("%v\n", *emp4)
	fmt.Printf("%v\n", emp4)
}
