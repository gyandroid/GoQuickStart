package main

import "fmt"

func main() {
	// Arrays - Go have to be of fixed length and type
	// Slices - Array that doesn't have fixed type

	//var fruitArr [2]string

	// Assign value
	//fruitArr[0] = "Apple"
	//fruitArr[1] = "ORange"

	// Declare and Assign
	/*fruitArr := [2]string{"Apple", "ORange"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])*/

	fruitSlice := []string{"Apple", "ORange", "Grape"}
	fmt.Println(fruitSlice)
	fmt.Println(fruitSlice[2])
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:2])
}