package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p *person) updateFirstName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
