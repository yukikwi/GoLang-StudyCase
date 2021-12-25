package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello world\n")
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome to index\n")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)

	fmt.Println("Start http server on port 8080")
}
