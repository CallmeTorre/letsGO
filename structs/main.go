package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	//contact   contactInfo
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateFirstName(newFirstName string) {
	p.firstName = newFirstName
}

func main() {
	var alex person = person{
		firstName: "Alex",
		lastName:  "Anderson",
		/*contact: contactInfo{
			email: "a@a.com",
			zip:   53500,
		},*/
		contactInfo: contactInfo{
			email: "a@a.com",
			zip:   53500,
		},
	}
	alex.print()
	alex.updateFirstName("Alexi")
	alex.print()
}
