package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	// can write 2 ways
	// first
	// var a, b string
	// a, b = swap("hello", "world")

	// or
	a, b := swap("hello", "world")
	// := mean declartion and assign
	// =  mean assign value only
	fmt.Println(a, b)
}
