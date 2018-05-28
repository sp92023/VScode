package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

/*type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}*/

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	/*alex := person{"Alex", "Anderson"} // {Alex Anderson}
	alex2 := person{firstName: "Alex", lastName: "Anderson"} // {Alex Anderson}
	alex3 := person{lastName: "Anderson", firstName: "Alex"} // {Alex Anderson}
	fmt.Println(alex)
	fmt.Println(alex2)
	fmt.Println(alex3)*/

	/* **zero value
	string:""
	int:0
	float:0
	bool:false*/
	/*var alex person

	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	fmt.Println(alex)
	fmt.Printf("%+v", alex)*/

	/*jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}*/
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	/*//fmt.Printf("%+v", jim)
	jim.updateName("jimmy")
	jim.print()*/

	/*jimPointer := &jim //operator; jim is a reference to the struct in memory
	// use the &, we turn &jim into a memory address or a pointer
	// and then assign that value to jimPointer
	jimPointer.updateName("jimmy")
	jim.print()*/

	jim.updateName("jimmy")
	jim.print()
	// jimPointer -> type of *person; jim -> type of person
	// '*person' -> type of a pointer to a person
	// Go will automatically turn variable of type person into pointer person for you
}

/*func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}*/

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// '*person' is a type description; a pointer that points at a person
// '*pointerToPerson' is an operator; it means we want to manipulate the value the pointer is referencing

// turn address into value with '*address'
// turn value into address with '&value'

func (p person) print() {
	fmt.Printf("%+v", p)
}

// value types: use pointers to change these things in a function
// EX. int, float, string, bool, structs

// reference types: don't worry about pointers with these
// EX. slices, maps, channels, pointers, functions
