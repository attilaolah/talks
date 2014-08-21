package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"strings"
)

// START2 OMIT
func fib(n int) *big.Int {
	a, b := big.NewInt(0), big.NewInt(1)
	for ; n != 0; n-- {
		a, b = b, a.Add(a, b)
	}
	return a
}

func fibn(from, to int) []*big.Int { // HL
	result := []*big.Int{}    // HL
	for ; from < to; from++ { // HL
		result = append(result, fib(from)) // HL
	} // HL
	return result // HL
} // HL

// END2 OMIT

func main() {
	http.HandleFunc("/fib/", func(w http.ResponseWriter, r *http.Request) {
		// START1 OMIT
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
			json.NewEncoder(w).Encode(fibn(from, to)) // HL
			return
		}
		// END1 OMIT
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
