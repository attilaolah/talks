package main

import (
	"fmt"
	"math/big"
)

// fib returns a function that returns successive Fibonacci numbers.
func fib() func() *big.Int {
	a, b := big.NewInt(0), big.NewInt(1)
	return func() *big.Int {
		a, b = b, a.Add(a, b)
		return a
	}
}

func main() {
	f := fib()
	for {
		fmt.Println(f().String())
	}
}
