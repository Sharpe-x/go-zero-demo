package example

import (
	"fmt"
	"strconv"
)

func TestRange() {
	var employees []*Employee
	for i := 0; i < 10; i++ {
		employees = append(employees, &Employee{
			Name: strconv.Itoa(i),
		})
	}

	mapNameToEmployees := make(map[string]*Employee)
	for _, employee := range employees {
		fmt.Println(employee)
		mapNameToEmployees[employee.Name] = employee
	}

	fmt.Println("-----------------------------------")
	for _, employee := range mapNameToEmployees {
		fmt.Println(employee)
	}
}
