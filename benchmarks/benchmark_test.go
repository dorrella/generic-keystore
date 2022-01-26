package main

import "testing"

func BenchmarkIntReads(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkIntReads(false)
	}
}

func BenchmarkIntWrites(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkIntWrites(false)
	}
}

func BenchmarkStringReads(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkStringReads(false)
	}
}

func BenchmarkStringWrites(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkStringWrites(false)
	}
}
