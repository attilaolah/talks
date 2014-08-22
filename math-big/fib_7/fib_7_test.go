package main

import "testing"

type sink struct{}

func (s *sink) Write(b []byte) (int, error) {
	return len(b), nil
}

func BenchmarkFib12_18(b *testing.B) {
	s := &sink{}
	for i := 0; i < b.N; i++ {
		fibn(s, 12, 18)
	}
}

func BenchmarkFib1234_5678(b *testing.B) {
	s := &sink{}
	for i := 0; i < b.N; i++ {
		fibn(s, 1234, 5678)
	}
}
