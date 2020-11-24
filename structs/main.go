package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	logan := person{
		firstName: "Logan",
		lastName:  "Wood",
		contact: contactInfo{
			email:   "lwood3499@gmail.com",
			zipCode: 12345,
		},
	}

	// option 1: explicitly defining pointer to memory, then use that pointer as
	// the receiver in a function to change a property of the value at the memory location
	// loganPointer := &logan
	// loganPointer.updateFirstName("Stretch")

	// option 2: shortcut
	// NOTE: Go is automatically defining our pointer within the method with this shortcut
	// as long as the '*' in *person is used, if the '*' is left off, obviously it is just the normal type
	// and it won't work to get the value at that address
	logan.updateFirstName("Stretch")
	logan.print()
}

// NOTE: *person is describing the type to be used as a "pointer to person type"
// the '*' in *pointerToPerson is an operator getting the value at the address provided by the pointer being
// passed in the parameter
func (pointerToPerson *person) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// NOTE: You can...
// turn an 'address' into a 'value' with *address
// turn a 'value' into an 'address' with &value

func (p person) print() {
	fmt.Printf("%+v", p)
}

// NOTE: different types make use of of underlying data strctures already and therefore
// don't need the use of pointers to make changes. Here is a table
////     Value Types (use pointers to change these in a function)    //    Reference Types (don't need pointers to change) ////
//      - int                                                        //    - slices
//      - float                                                      //    - maps
//      - string                                                     //    - channels
//      - bool                                                       //    - pointers
//      - structs                                                    //    - functions
