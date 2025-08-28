package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (e *Employee) GetDetails() string {
	return fmt.Sprintf("Name: %s, Age: %d, EmployeeID: %s", e.name, e.age, e.EmployeeID)
}

func main() {
	e := Employee{
		Person:     Person{name: "Alice", age: 30},
		EmployeeID: "E12345",
	}
	fmt.Println(e.GetDetails())
}
