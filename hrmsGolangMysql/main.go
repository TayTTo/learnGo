package main

import (
	"fmt";
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Hello")
}

func main() {
	fmt.Println("Running on port 8080")
	http.HandleFunc("/", d)
	http.ListenAndServe(":8080", nil)
}
