package main

import (
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strconv"
	"strings"
)

func fib(n int) *big.Int {
	a, b := big.NewInt(0), big.NewInt(1)
	for ; n != 0; n-- {
		a, b = b, a.Add(a, b)
	}
	return a
}

// START OMIT
func fibn(w io.Writer, from, to int) { // HL
	if from >= to { // HL
		w.Write([]byte{'[', ']'}) // HL
		return                    // HL
	} // HL
	w.Write([]byte{'['}) // HL
	a, b := big.NewInt(0), big.NewInt(1)
	for i := from; i != 0; i-- {
		a, b = b, a.Add(a, b)
	}
	for i := to - from - 1; i != 0; i-- {
		fmt.Fprintf(w, "%s,", a) // HL
		a, b = b, a.Add(a, b)
	}
	fmt.Fprintf(w, "%s", a) // HL
	w.Write([]byte{']'})    // HL
}

// END OMIT

func main() {
	http.HandleFunc("/fib/", func(w http.ResponseWriter, r *http.Request) {
		parts := strings.SplitN(r.URL.Path[len("/fib/"):], "..", 2)
		if len(parts) == 2 {
			// /fib/from..to
			from, err := strconv.Atoi(parts[0])
			if err != nil {
				http.NotFound(w, r)
				return
			}
			to, err := strconv.Atoi(parts[1])
			if err != nil {
				http.NotFound(w, r)
				return
			}
			fibn(w, from, to)
			return
		}
		if len(parts) == 1 {
			// /fib/n
			if n, err := strconv.Atoi(parts[0]); err == nil {
				fmt.Fprintf(w, "%d", fib(n))
				return
			}
			http.NotFound(w, r)
			return
		}
		http.NotFound(w, r)
	})
	http.ListenAndServe(":8080", nil)
}
