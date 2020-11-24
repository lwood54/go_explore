package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// option 1 to declare a variable with struct type
	// logan := person{"Logan", "Wood"} // not ideal because order of struct may change, messing definition up
	// option 2
	logan := person{firstName: "Logan", lastName: "Wood"} // more explicit (and familiar)
	fmt.Println("logan: ", logan)
	fmt.Println("firstName: ", logan.firstName)
	fmt.Printf("%+v", logan) // %+v results in printing fieldnames and values

	// option 3
	var tiffany person
	fmt.Println("Tiffany: ", tiffany)
	tiffany.firstName = "Tiffany"
	tiffany.lastName = "Wood"
	fmt.Printf("%v", tiffany)
}
