package bench_test

import (
	"fmt"
	"math"
	"reflect"
	"testing"

	"github.com/shamaton/msgpack"
	aaaa "github.com/vmihailenco/msgpack"
)

func BenchmarkA1(b *testing.B) {
	v := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		v[i] = i
	}
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			v[j]++
		}
	}
}

func Benchmark2(b *testing.B) {
	v := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		v[i] = i
	}
	for i := 0; i < b.N; i++ {
		for _, vv := range v {
			vv++
		}
	}
}

type st struct {
	a int
	b *uint
}

func BenchmarkShamaton(b *testing.B) {
	//v := []int{1, 2, 3, math.MinInt64}
	/*
		v := make([]int, 10000)
		for i := 0; i < 10000; i++ {
			v[i] = i
		}
	*/
	// v := 777
	//v := "thit is test"
	v := st{a: math.MinInt32, b: nil}
	/*
		v := map[int]map[int]int{}
		for i := 0; i < 10000; i++ {
			v[i] = map[int]int{}
			for j := 0; j < 10; j++ {
				v[i][j] = i * j
			}
		}
	*/
	//v := time.Now()
	for i := 0; i < b.N; i++ {
		_, err := msgpack.SerializeAsMap(v)
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
	// v := 777
	//v := "thit is test"
	v := st{a: math.MinInt32, b: nil}
	/*
		v := map[int]map[int]int{}
		for i := 0; i < 10000; i++ {
			v[i] = map[int]int{}
			for j := 0; j < 10; j++ {
				v[i][j] = i * j
			}
		}
	*/
	//v := time.Now()

	for i := 0; i < b.N; i++ {
		_, err := aaaa.Marshal(v)
		/*
			var buf bytes.Buffer
			enc := aaaa.NewEncoder(&buf).StructAsArray(true)
			err := enc.Encode(v)
		*/
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

func bench_c(v interface{}) {
	switch v := v.(type) {
	case int8:
		v++
	case int16:
		v++
	case int32:
		v++
	case int:
		v++
	case int64:
		v++
	}
}

func bench_d(rv reflect.Value) {
	switch rv.Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		v := rv.Int()
		v++
	}
}
