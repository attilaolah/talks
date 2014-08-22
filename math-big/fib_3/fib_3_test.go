package main

import (
	"encoding/json"
	"testing"
)

type sink struct{}

func (s *sink) Write(b []byte) (int, error) {
	return len(b), nil
}

func BenchmarkFib12_18(b *testing.B) {
	e := json.NewEncoder(&sink{})
	for i := 0; i < b.N; i++ {
		e.Encode(fibn(12, 18))
	}
}

func BenchmarkFib1234_4568(b *testing.B) {
	e := json.NewEncoder(&sink{})
	for i := 0; i < b.N; i++ {
		e.Encode(fibn(1234, 4568))
	}
}
