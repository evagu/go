package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")

}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Server starting")
	http.ListenAndServe(":300", nil)
}
