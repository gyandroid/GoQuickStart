package main

import "fmt"

func main() {
	// Define map
	/*emails := make(map[string]string)

	// Assign key, values
	emails["Bob"] = "bob@gmail.com"
	emails["Alice"] = "alice@gmail.com"
	emails["Mike"] = "mike@gmail.com"*/

	// declare map and add kv
	emails := map[string]string{"Bob":"bob@gmail.com", "Alice":"alice@gmail.com"}
	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	// delete friom map
	delete(emails, "Bob")

	fmt.Println(emails)
}