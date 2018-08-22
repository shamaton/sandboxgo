package bench_test

import (
	"fmt"
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
	Map    map[string]uint
	Child  BenchChild
}

var vv = BenchMarkStruct{
	iInt:   -123,
	Uint:   456,
	Float:  1.234,
	Double: 6.789,
	Bool:   true,
	String: "this is text.",
	Array:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	Map:    map[string]uint{"this": 1, "is": 2, "map": 3},
	Child:  BenchChild{Int: 123456, String: "this is struct of child"},
}

// var v = []uint{1, 2, 3, 4, 5, 6, math.MaxUint64}
var v = []string{"this", "is", "test"}

//var v = [4]string{"this", "is", "test"}

// var v = []float32{1.23, 4.56, math.MaxFloat32}
// var v = []float64{1.23, 4.56, math.MaxFloat64}
// var v = []bool{true, false, true}
// var v = []uint8{0x82, 0xa1, 0x41, 0x07, 0xa1, 0x42, 0xa1, 0x37}
// var v = [8]uint8{0x82, 0xa1, 0x41, 0x07, 0xa1, 0x42, 0xa1, 0x37}

// var v = map[string]BenchMarkStruct{"a": vv, "b": vv}
// var v = map[string]float32{"1": 2.34, "5": 6.78}
// var v = map[string]bool{"a": true, "b": false}

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
	// v := st{A: math.MinInt32, B: math.MaxUint64, C: -1}
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
	// v := st{A: math.MinInt32, B: math.MaxUint64, C: -1}
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

var data []byte
var e2 error

func init() {
	v := map[string]int32{"a": 1, "b": 2, "c": 3}
	data, e2 = msgpack.SerializeAsArray(v)
	if e2 != nil {
		fmt.Println("init err : ", e2)
	}
}

func BenchmarkDesShamaton(b *testing.B) {
	var r map[string]int
	for i := 0; i < b.N; i++ {
		err := msgpack.Deserialize(data, &r)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
func BenchmarkDesVmihailenco(b *testing.B) {
	var r map[string]int
	for i := 0; i < b.N; i++ {
		err := aaaa.Unmarshal(data, &r)
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
