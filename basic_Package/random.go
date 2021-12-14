package main

import "fmt"
import "math/rand"


func main() {
	fmt.Println("My favorite number is", rand.Intn(100))
	fmt.Println("My favorite number is", rand.Intn(100))
	// output is same on every run time due to seed
}
