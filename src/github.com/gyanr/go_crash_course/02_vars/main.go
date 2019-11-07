package main

import "fmt"

//var name  = "Gyan"


func main() {
	// var
	// const - similar to final
	// shorthand method to define variable ":=" -- to be used inside a function
	// var name string = "Gyan"
	// var age int = 33
	//var name  = "Gyan"
	var age int32  = 33
	var isCool = true
	isCool = false

	// Shorthand
	//name := "Gyan"
	size := 1.5
	//email := "gyan@gmail.com"

	name, email := "Gyan", "gyan@gmail.com"

	fmt.Println(name, age, isCool, size, email)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", size)
}