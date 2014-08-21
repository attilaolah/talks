package main

import "testing"

func BenchmarkFib12_18(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibn(12, 18)
	}
}
