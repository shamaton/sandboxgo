package bench_test

import (
	"fmt"
	"math"
	"reflect"
	"testing"

	"github.com/shamaton/msgpack"
	aaaa "github.com/vmihailenco/msgpack"
)

type st struct {
	A int
	B uint64
	C int
}

/*
type BenchChild struct {
	Int    int
	String string
}
type BenchMarkStruct struct {
	iInt   int
	Uint   uint
	Float  float32
	Double float64
	Bool   bool
	String string
	Array  []int
	Map    map[string]int
	Child  BenchChild
}

var v = BenchMarkStruct{
	iInt:   -123,
	Uint:   456,
	Float:  1.234,
	Double: 6.789,
	Bool:   true,
	String: "this is text.",
	Array:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	Map:    map[string]int{"this": 1, "is": 2, "map": 3},
	Child:  BenchChild{Int: 123456, String: "this is struct of child"},
}
*/

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
	v := st{A: math.MinInt32, B: math.MaxUint64, C: -1}
	//v := map[int]interface{}{1: 2, 3: "a", 4: []float32{1.23}}
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
		//_, err := msgpack.SerializeAsArray(v)
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
	v := st{A: math.MinInt32, B: math.MaxUint64, C: -1}
	//v := map[int]interface{}{1: 2, 3: "a", 4: []float32{1.23}}
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
