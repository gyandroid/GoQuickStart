package main

import "fmt"

func main() {

	x := 15
	y :=10
	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// else-if

	color := "BLUE"
	if color == "red" {
		fmt.Println("Color is red")
	} else if color == "blue" || color == "BLUE" {
		fmt.Println("Color is blue")
	} else {
		fmt.Println("Color is NOT blue or RED")
	}

	// Switch

	switch color {
	case "red" :
		fmt.Println("Color is red")
	case "blue":
	case "BLUE":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is NOT blue or RED")
	}
}