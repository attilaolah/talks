package main

// STARTI OMIT
import (
	big "github.com/ncw/gmp"
)

// ENDI OMIT

// START1 OMIT
func bad(a, b *big.Int) *big.Int {
	z := big.NewInt(0)
	z.Add(a, b)
	return z
}

// END1 OMIT

// START2 OMIT
func good(a, b *big.Int) *big.Int {
	return a.Add(a, b)
}

// END2 OMIT
