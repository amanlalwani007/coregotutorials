package main

import (
	"fmt"
	"strconv"
)

//methods are of two types
//value receivers and  pointers receivers
//Define struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int
	firstName, lastName, city, gender string
	age                               int
}

func (p Person) greet() string {
	return "Hello,my name is " + p.firstName + " " + p.lastName + " " + strconv.Itoa(p.age)
}

//has birthday method(pointer receiver)

func (p *Person) hasbirthday() {
	p.age++
}

func main() {
	person1 := Person{firstName: "aman", lastName: "lalwani", city: "agra", age: 20, gender: "M"}
	fmt.Println(person1)
	fmt.Println(person1.firstName)
	person1.hasbirthday()
	fmt.Println(person1.greet())

}
