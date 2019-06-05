package main

import "fmt"

func main() {
	name, email := "Brad", "brad@gmail.com"
	var age = 37
	var isCool = true
	isCool = false
	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", isCool)
}
