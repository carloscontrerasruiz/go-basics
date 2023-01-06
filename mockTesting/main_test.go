package main

import "testing"

func TestGetFulltimeEmployeeById(t *testing.T) {
	table := []struct {
		id          int
		dni         string
		mockFunc    func()
		expextedEmp FullTimeEmployee
	}{
		{
			id:  1,
			dni: "1",
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						Id:       1,
						Position: "CEO",
					}, nil
				}

				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{
						Name: "Jon Doe",
						Age:  35,
						Dni:  "1",
					}, nil
				}
			},

			expextedEmp: FullTimeEmployee{
				Person: Person{
					Name: "Jon Doe",
					Age:  35,
					Dni:  "1",
				},
				Employee: Employee{
					Id:       1,
					Position: "CEO",
				},
			},
		},
	}

	originalGetEmployee := GetEmployeeById
	originalGetPersonDni := GetPersonByDNI

	for _, test := range table {
		test.mockFunc()
		ft, err := GetFulltimeEmployeeById(test.id, test.dni)

		if err != nil {
			t.Errorf("Error getting the employee")
		}

		if ft.Age != test.expextedEmp.Age {
			t.Errorf("Las edades no coinciden")
		}
	}

	GetEmployeeById = originalGetEmployee
	GetPersonByDNI = originalGetPersonDni
}
