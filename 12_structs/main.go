package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastNume  string
	city      string
	gender    string
	age       int
}

func (p Person) greet() string {
	return "Hello " + p.firstName + " " + p.lastNume + " " + strconv.Itoa(p.age)
}

func (p *Person) hasBirthDay() {
	p.age++
}
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastNume = spouseLastName
	}
}

func main() {
	person1 := Person{firstName: "Samen", lastNume: "Simith", city: "Boston", gender: "f", age: 34}
	// fmt.Println(person1.firstName)
	person1.hasBirthDay()
	person1.getMarried("Wili")
	fmt.Println(person1.greet())

}
