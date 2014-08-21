package main

import "testing"

func BenchmarkFib37(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(37)
	}
}

func BenchmarkFib93(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(93)
	}
}
