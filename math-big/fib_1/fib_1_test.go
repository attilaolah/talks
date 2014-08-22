package main

import (
	"fmt"
	"testing"
)

type sink struct{}

func (s *sink) Write(b []byte) (int, error) {
	return len(b), nil
}

func BenchmarkFib37(b *testing.B) {
	s := &sink{}
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(s, "%d", fib(37))
	}
}
