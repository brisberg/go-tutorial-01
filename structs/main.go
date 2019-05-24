package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "alex@gmail.com",
			zipCode: 54321,
		},
	}

	alex.updateName("Alexander")
	alex.print()
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p *person) print() {
	fmt.Printf("%+v", p)
}
