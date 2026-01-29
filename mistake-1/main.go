package main

import (
	"fmt"
)

// using nil slice and maps
func main() {
	// here map is delcare but didn't intialize it
	var a map[string]int

	// this is correct behaviour
	// var a make(map[string]int)
	a["any"] = 1

	fmt.Println("hello")
}
