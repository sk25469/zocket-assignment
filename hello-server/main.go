package main

import (
	"fmt"
	"net/http"
)

func helloFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", helloFunc)
	http.ListenAndServe(":3000", nil)
}
