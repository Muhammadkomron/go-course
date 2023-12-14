package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	//alex := person{
	//	firstName: "Alex",
	//	lastName:  "Anderson",
	//}
	//alex := person{"Alex", "Anderson"}
	//var alex person
	//alex.firstName = "Alex"
	//alex.lastName = "Anderson"
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "muhammadkomron.qobilov@gmail.com",
			zipCode: 100069,
		},
	}
	//alexPointer := &alex
	//alexPointer.updateFirstName("Alexander")
	alex.updateFirstName("Alexander")
	alex.print()
	fmt.Println(*&alex)
}
