package main

import (
	"errors"
	"fmt"
	"log"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func main() {

	employee, err := getInformation(1001)
	if err != nil {
		log.Panicln(err)
	} else {
		log.Printf("%v", employee)
	}
	fmt.Println("next")
}

var ErrNotFound = errors.New("Employee not found!")

func getInformation(id int) (e *Employee, err error) {

	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
		}

	}()
	e, err = apiCallEmployee(1000)
	panic("there is a panic")
	return e, err
}

func apiCallEmployee(id int) (*Employee, error) {
	employee := Employee{LastName: "Doe", FirstName: "John"}
	return &employee, nil
}
