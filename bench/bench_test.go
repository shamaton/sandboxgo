package bench_test

import (
	"testing"
)

func BenchmarkA1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f(i)
	}
}

func Benchmark2(b *testing.B) {
	f = bench_a
	for i := 0; i < b.N; i++ {
		f(i)
	}
}

var f = bench_a

func bench_a(v int) int {
	v++
	return v
}
func bench_b(v int) int {
	for i := 0; i < 1000; i++ {

		v++
	}
	return v
}
