package main

import "fmt"

func main() {

	ids :=[]int{33,25,342,45} // slice

	// Loops through ids
	for i,id := range ids {
		fmt.Printf("%d - ID: %d \n", i, id)
	}

	// not using index

	for _,id := range ids {
		fmt.Printf("ID: %d \n", id)
	}

	// Add ids together
	sum := 0
	for _,id := range ids {
		sum += id
	}
	fmt.Printf("Sum: %d \n", sum)

	// Range with maps
	emails := map[string]string{"Bob":"bob@gmail.com", "Alice":"alice@gmail.com"}

	for k,v := range emails {
		fmt.Printf("%s:%s \n", k, v)
	}
	for k := range emails {
		fmt.Printf("Name: %s \n", k)
	}
}