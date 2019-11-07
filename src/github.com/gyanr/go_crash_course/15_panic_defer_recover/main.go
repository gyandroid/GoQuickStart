package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
)

func main() {
	// Initial example
	/*fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")*/

	// HTTP example

	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)

}