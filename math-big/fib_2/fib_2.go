package main

import (
	"fmt"
	"math/big"
	"net/http"
	"strconv"
)

// START OMIT
func fib(n int) *big.Int { // HL
	a, b := big.NewInt(0), big.NewInt(1) // HL
	for ; n != 0; n-- {
		a, b = b, a.Add(a, b) // HL
	}
	return a
}

// END OMIT

func main() {
	http.HandleFunc("/fib/", func(w http.ResponseWriter, r *http.Request) {
		if n, err := strconv.Atoi(r.URL.Path[len("/fib/"):]); err == nil {
			fmt.Fprintf(w, "%d", fib(n))
			return
		}
		http.NotFound(w, r)
	})
	http.ListenAndServe(":8080", nil)
}
