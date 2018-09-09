package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math"
	"reflect"
	"time"

	"github.com/shamaton/msgpack"
	exttime "github.com/shamaton/msgpack/time"
	aaaa "github.com/vmihailenco/msgpack"
)

// allocate memo
// make
// struct -> rv.Type().Field(i)
// rv.Interface().(type) deserliaze only
// rv.Set(slice)

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
	Array  *[]int
	Map    map[string]int
	Child  BenchChild
}

var vv = BenchMarkStruct{
	iInt:   -123,
	Uint:   456,
	Float:  1.234,
	Double: 6.789,
	Bool:   true,
	String: "this is text.",
	Array:  &[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	Map:    map[string]int{"this": 1, "is": 2, "map": 3},
	Child:  BenchChild{Int: 123456, String: "this is struct of child"},
}

func init() {
	fmt.Println("set ext func")
	msgpack.AddExtCoder(exttime.Encoder, exttime.Decoder)
	msgpack.RemoveExtCoder(exttime.Encoder, exttime.Decoder)
}

type benchmarkStruct struct {
	Name   string
	Age    int
	Colors []string
	Data   []byte
	//CreatedAt time.Time
	//UpdatedAt time.Time
}

func structForBenchmark() *benchmarkStruct {
	return &benchmarkStruct{
		Name:   "Hello World",
		Age:    math.MaxInt32,
		Colors: []string{"red", "orange", "yellow", "green", "blue", "violet"},
		Data:   make([]byte, 10),
		//CreatedAt: time.Now(),
		//UpdatedAt: time.Now(),
	}
}

type benchmarkStructPartially struct {
	Name string
	Age  int
}

func main() {
	type stt struct {
		Int int
	}
	type stt2 struct {
		Int interface{}
	}
	v := stt{Int: 1}
	fmt.Println(v)
	//v := map[interface{}]int{"a": 1, 6666: 2, "c": 3}
	var vr, vr2 stt2
	var sr, sr2 stt2

	fmt.Println("vv :", vv)
	for i := 0; i < 100; i++ {
		var r BenchMarkStruct
		dd, _ := msgpack.EncodeStructAsArray(vv)
		err := msgpack.DecodeStructAsArray(dd, &r)
		if err != nil {
			fmt.Println("des err : ", err)
			return
		}
		if i == 0 || i == 99 {
			fmt.Println(i, r)
		}
	}
	fmt.Println("-------------------vmi arr-----------------------")
	{
		d := vmiMarshalArray(v)
		fmt.Println(hex.Dump(d))

		err := aaaa.Unmarshal(d, &vr)
		if err != nil {
			fmt.Println("des err : ", err)
		}
		fmt.Println(vr, reflect.ValueOf(vr).Type())

		d2 := vmiMarshalArray(vr)
		fmt.Println(hex.Dump(d2))
	}
	fmt.Println("-------------------sha arr-----------------------")
	{
		d, _ := shamaton(v)
		fmt.Println(hex.Dump(d))
		err := msgpack.DecodeStructAsArray(d, &sr)
		if err != nil {
			fmt.Println("des err : ", err)
		}
		fmt.Println(sr, reflect.ValueOf(sr).Type())
		d2, _ := msgpack.EncodeStructAsArray(sr)
		fmt.Println(hex.Dump(d2))
	}

	fmt.Println("-------------------vmi map-----------------------")
	{
		d := vmiMarshalMap(v)
		fmt.Println(hex.Dump(d))

		err := aaaa.Unmarshal(d, &vr2)
		if err != nil {
			fmt.Println("des err : ", err)
		}
		fmt.Println(vr2)

		d2 := vmiMarshalMap(vr2)
		fmt.Println(hex.Dump(d2))
	}

	fmt.Println("-------------------sha map-----------------------")
	{
		_, d := shamaton(v)
		fmt.Println(hex.Dump(d))

		err := msgpack.DecodeStructAsMap(d, &sr2)
		if err != nil {
			fmt.Println("des err : ", err)
		}
		fmt.Println(sr2)
		d2, _ := msgpack.EncodeStructAsMap(sr2)
		fmt.Println(hex.Dump(d2))
	}

	var vvvv = math.MaxInt32 - 1
	var rrrr uint32
	tttt(vvvv, &rrrr)
	fmt.Println("vvvv : ", vvvv, " | rrrr : ", rrrr)
}

func tttt(v interface{}, r interface{}) {
	rv := reflect.ValueOf(v)
	rvv := reflect.ValueOf(r)
	rvv = rvv.Elem()
	i := rv.Int()
	rvv.SetUint(uint64(i))
}

func _main() {
	type st struct {
		A int
		b *uint
		c int
	}

	// a.F()
	//v := []int{1, 2, 3, math.MinInt64}
	//v := [5]int{1, 2, 3, math.MinInt64}
	//v = nil
	//v := "this is test"
	//v := []bool{true, false}
	// v := float64(math.MaxFloat64)
	// v := []byte{0x82, 0xa1, 0x41, 0x07, 0xa1, 0x42, 0xa1, 0x37}
	//v := []uint8{0x82, 0xa1, 0x41, 0x07, 0xa1, 0x42, 0xa1, 0x37}
	// v := [8]uint8{0x82, 0xa1, 0x41, 0x07, 0xa1, 0x42, 0xa1, 0x37}
	// v := &st{A: math.MinInt32, b: nil}
	//v := map[int]interface{}{1: 2, 3: "a", 4: []float32{1.23}}
	v := time.Now()
	// v := float32(1.234)
	// v := map[string]float64{"1": 2.34, "5": 6.78}
	// v := map[string]bool{"a": true, "b": false}
	//v := []bool{true, false, true}
	vd1, vd2 := vmihailenco(v)
	sd1, sd2 := shamaton(v)
	fmt.Println("vmihaile arr : ", hex.Dump(vd1))
	fmt.Println("shamaton arr : ", hex.Dump(sd1))
	fmt.Println("vmihaile map : ", hex.Dump(vd2))
	fmt.Println("shamaton map : ", hex.Dump(sd2))
}

func shamaton(v interface{}) ([]byte, []byte) {
	d, err := msgpack.EncodeStructAsArray(v)
	if err != nil {
		fmt.Println("err arr : ", err)
	}
	d2, err := msgpack.EncodeStructAsMap(v)
	if err != nil {
		fmt.Println("err map : ", err)
	}
	return d, d2
}

func vmihailenco(v interface{}) ([]byte, []byte) {
	d1 := vmiMarshalArray(v)
	d2 := vmiMarshalMap(v)
	return d1, d2
}

func vmiMarshalArray(v interface{}) []byte {
	var buf bytes.Buffer
	enc := aaaa.NewEncoder(&buf).StructAsArray(true)
	err := enc.Encode(v)
	if err != nil {
		fmt.Println("err arr : ", err)
	}
	return buf.Bytes()
}

func vmiMarshalMap(v interface{}) []byte {
	d, err := aaaa.Marshal(v)
	if err != nil {
		fmt.Println("err map : ", err)
	}
	return d
}
