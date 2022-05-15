package main

import "testing"

// go test -bench=. bench_test.go
// go test -bench . -benchmem single_number_bench_test.go
// go test -bench . -benchmem -cpuprofile=cpu.out -memprofile=mem.out -memprofilerate=1 pool_test.go

func BenchmarkSingleNumberXOR(b *testing.B) {
	for i := 0; i < b.N; i++ {
		singleNumberXOR([]int{4, 1, 2, 1, 2})
	}
}

func BenchmarkSingleNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		singleNumberXOR([]int{4, 1, 2, 1, 2})
	}
}
