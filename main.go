package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	jim.updateName("Jimmy")
	jim.print()
}

func (p *person) updateName(firstName string) {
	(*p).firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
