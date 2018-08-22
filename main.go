package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math"
	"reflect"

	"github.com/shamaton/msgpack"
	a "github.com/shamaton/sandboxgo/msgpack"
	aaaa "github.com/vmihailenco/msgpack"
)

// allocate memo
// make
// struct -> rv.Type().Field(i)
// rv.Interface().(type) deserliaze only
// rv.Set(slice)

func _main() {
	v := []int{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000}
	var vr [3]int
	var sr [3]int

	switch reflect.ValueOf(sr).Interface().(type) {
	case []int:
		fmt.Println("iiiiiiiiiiiiiii")
	case [len(sr)]int:
		fmt.Println("aaaaaaaaaaaaa")
	}

	d := vmiMarshalMap(v)
	fmt.Println(hex.Dump(d))

	err := aaaa.Unmarshal(d, &vr)
	if err != nil {
		fmt.Println("des err : ", err)
	}
	fmt.Println(vr)

	d2 := vmiMarshalMap(vr)
	fmt.Println(hex.Dump(d2))

	fmt.Println("------------------------------------------")

	d, _ = shamaton(v)
	fmt.Println(hex.Dump(d))
	err = msgpack.Deserialize(d, &sr)
	if err != nil {
		fmt.Println("des err : ", err)
	}
	fmt.Println(sr)
	d2, _ = msgpack.Serialize(sr)
	fmt.Println(hex.Dump(d2))

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

func main() {
	type st struct {
		A int
		b *uint
		c int
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
		Map    map[string]int
		Child  BenchChild
	}
	/*
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

	a.F()
	//v := []int{1, 2, 3, math.MinInt64}
	v := [5]int{1, 2, 3, math.MinInt64}
	//v = nil
	//v := "this is test"
	//v := []bool{true, false}
	// v := float64(math.MaxFloat64)
	// v := []byte{0x82, 0xa1, 0x41, 0x07, 0xa1, 0x42, 0xa1, 0x37}
	//v := []uint8{0x82, 0xa1, 0x41, 0x07, 0xa1, 0x42, 0xa1, 0x37}
	// v := [8]uint8{0x82, 0xa1, 0x41, 0x07, 0xa1, 0x42, 0xa1, 0x37}
	// v := &st{A: math.MinInt32, b: nil}
	//v := map[int]interface{}{1: 2, 3: "a", 4: []float32{1.23}}
	//v := time.Now()
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
	d, err := msgpack.SerializeAsArray(v)
	if err != nil {
		fmt.Println("err arr : ", err)
	}
	d2, err := msgpack.SerializeAsMap(v)
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

func vmiMarshalMap(v interface{}) []byte {
	var buf bytes.Buffer
	enc := aaaa.NewEncoder(&buf).StructAsArray(true)
	err := enc.Encode(v)
	if err != nil {
		fmt.Println("err arr : ", err)
	}
	return buf.Bytes()
}

func vmiMarshalArray(v interface{}) []byte {
	d, err := aaaa.Marshal(v)
	if err != nil {
		fmt.Println("err map : ", err)
	}
	return d
}
