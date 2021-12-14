package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("My favorite number is", rand.Intn(10))
	// output is same on every run time due to seed
}
