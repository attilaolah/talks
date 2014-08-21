package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// START OMIT
func fib(n int) int {
	a, b := 0, 1
	for ; n != 0; n-- {
		a, b = b, a+b
	}
	return a
}

func main() {
	http.HandleFunc("/fib/", func(w http.ResponseWriter, r *http.Request) {
		if n, err := strconv.Atoi(r.URL.Path[len("/fib/"):]); err == nil {
			fmt.Fprintf(w, "%d", fib(n)) // HL
			return
		}
		http.NotFound(w, r)
	})
	http.ListenAndServe(":8080", nil)
}

// END OMIT
