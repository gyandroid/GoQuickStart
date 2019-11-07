package main

import (	"strconv"
	"fmt"
)

// Define person struct
type Person struct {
	firstName string
	lastName string
	city string
	gender string
	age int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello this is : "+p.firstName + " "+p.lastName+" my age is "+strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried method (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// init Person using struct
	person1 := Person{"Gavin", "Newsom", "SF", "M", 45}
	person2 := Person{"Samantha", "Newsom", "SF", "F", 45}
	/*fmt.Println(person1)
	fmt.Println(person1.lastName, person1.age)
	person1.age++
	person1.age++
	fmt.Println(person1)*/
	person1.hasBirthday()
	fmt.Println(person1.greet())
	person1.hasBirthday()
	person1.hasBirthday()
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
	person2.getMarried("Trump")
	fmt.Println(person2.greet())
}