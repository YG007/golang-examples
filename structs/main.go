package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	contactInfo //if no variable name is given, struct name is considered as variable name
}

func main() {

	//first way to assign a value to struct
	person1 := person{"Alex", "Anderson", contactInfo{}}
	person1.print()

	//second way to assign a value to struct
	person2 := person{firstName: "Alex", lastName: "Anderson"}
	person2.print()

	//third way to assign a value to struct
	var person3 person

	person3.firstName = "Alex"
	person3.lastName = "Anderson"
	person3.print()

	// working with embedded structs
	person4 := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@email.com",
			zipCode: 94000,
		},
	}
	person4.print()

	// creating pointers to structs in order to overcome pass by value contrsaints
	person4Pointer := &person4
	person4Pointer.updateName("Jimmy")
	person4.print()

	person5 := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@email.com",
			zipCode: 94000,
		},
	}
	person5.print()

	// shortcut in pointers, can call pointer receiver methods with the type as well
	person5.updateName("Jimmy")
	person5.print()

}

// methods with struct pointers as receivers
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// methods with structs as recerivers
func (p person) print() {
	fmt.Printf("%+v \n", p)
}
