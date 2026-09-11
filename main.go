package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Go application deployed using Argo CD!")
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Go application listening on port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}