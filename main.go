package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	result := addValue(300, 4)
	fmt.Fprintf(w, "Home: Added value = %d\n", result)
}
func About(w http.ResponseWriter, r *http.Request) {
	result := addValue(2, 4)
	fmt.Fprintf(w, "About: Added value = %d\n", result)
}

func addValue(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting server at port %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
