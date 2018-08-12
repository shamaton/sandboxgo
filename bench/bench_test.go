package bench_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/shamaton/msgpack"
	aaaa "github.com/vmihailenco/msgpack"
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

func BenchmarkShamaton(b *testing.B) {
	//v := []int{1, 2, 3, math.MinInt64}
	/*
		v := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			v[i] = i
		}
	*/
	v := 777
	for i := 0; i < b.N; i++ {
		_, err := msgpack.SerializeAsArray(v)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func BenchmarkVmihailenco(b *testing.B) {
	//v := []int{1, 2, 3, math.MinInt64}
	/*
		v := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			v[i] = i
		}
	*/
	v := 777
	var buf bytes.Buffer
	enc := aaaa.NewEncoder(&buf).StructAsArray(true)
	for i := 0; i < b.N; i++ {
		err := enc.Encode(v)
		if err != nil {
			fmt.Println(err)
			break
		}
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
