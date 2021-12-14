package main

import "fmt"

// naked return use only with short function
func swap(x, y string) (a, b string) {
	// doing something
	a = y
	b = x

	// return  a, b
	return
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
