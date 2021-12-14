package main
import "fmt"

// more about why type of variable is at the end of variable name: https://go.dev/blog/declaration-syntax
// in go we can group type of variables by declare type only a time at the last variable name
func add (a, b int) int{
	return a+b
}

func main() {
	fmt.Println(add(1, 2))
}