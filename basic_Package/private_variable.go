package main

import (
	"fmt"
	"math"
)

func main() {
	// in golang, if first variable name's character is lowercase so it is private variable. If it is uppercase so it is public variable that can access from outside
	
	// error: math.pi is private number
	// fmt.Println(math.pi)

	// this can access Pi variable under math package
	fmt.Println(math.Pi)
}
