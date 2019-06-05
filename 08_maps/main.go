package main

import "fmt"

func main() {
	emails := make(map[string]string)
	emails["Bob"] = "bob@gmail.com"
	emails["sheery"] = "she@gmail.com"

	delete(emails, "Bob")
	fmt.Println(emails)
}
