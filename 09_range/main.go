package main

import "fmt"

func main() {
	ids := []int{33, 76, 54, 23, 11, 2}
	for _, id := range ids {
		fmt.Printf("ID: %d \n", id)
	}
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("sum", sum)

	emails := make(map[string]string)
	emails["Bob"] = "bob@gmail.com"
	emails["sheery"] = "she@gmail.com"

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)

	}
}
