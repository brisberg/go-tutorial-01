package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{"Alex", "Anderson"} // Positional argument declaration
	barton := person{firstName: "Barton", lastName: "Beverly"}

	fmt.Println(alex)
	fmt.Println(barton)
}
