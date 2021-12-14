package main
import "fmt"

// more about why type of variable is at the end of variable name: https://go.dev/blog/declaration-syntax
func add (a int, b int) int{
	return a+b
}

func main() {
	fmt.Println(add(1, 2))
}