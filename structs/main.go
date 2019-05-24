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
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "alex@gmail.com",
			zipCode: 54321,
		},
	}

	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}
