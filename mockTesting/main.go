package main

import "time"

type Person struct {
	Name string
	Age  int
	Dni  string
}
type Employee struct {
	Id       int
	Position string
}
type FullTimeEmployee struct {
	Person
	Employee
}

func GetFulltimeEmployeeById(id int, dni string) (FullTimeEmployee, error) {
	var ftEmp FullTimeEmployee

	employee, err := GetEmployeeById(id)
	if err != nil {
		return ftEmp, err
	}

	ftEmp.Employee = employee

	person, err := GetPersonByDNI(dni)
	if err != nil {
		return ftEmp, err
	}

	ftEmp.Person = person

	return ftEmp, nil

}

var GetPersonByDNI = func(dni string) (Person, error) {
	//Select * from person where dni ....
	time.Sleep(5 * time.Second)
	return Person{}, nil
}

var GetEmployeeById = func(id int) (Employee, error) {
	//Select * from employees where dni ....
	time.Sleep(5 * time.Second)
	return Employee{}, nil
}
