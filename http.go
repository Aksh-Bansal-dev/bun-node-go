package main

import (
	"fmt"
	"net/http"
)

func main() {
	// http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, fib(40))
	})
	http.ListenAndServe(":4000", nil)
}

func fib(n int64) int64 {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
