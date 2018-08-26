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

func hogehoge() {
	var r interface{}
	r = ddd()
	rv := reflect.ValueOf(&r)
	fmt.Println("-------------------hogehogheoge-----------------------")
	fmt.Println(rv.Type())
	fmt.Println(rv.CanInterface(), rv.CanSet(), rv.Elem(), rv.Elem().Type())
	fmt.Println(rv.Elem().CanSet())

	switch r.(type) {
	case bool:
		fmt.Println("bool")
	case int64:
		fmt.Println("int644444444444!!")
	case int:
		fmt.Println("intttttttt!!")
	}

	switch rv.Elem().Kind() {
	case reflect.Interface:
		fmt.Println("kind inrefacade!!")
	}

	rvv := fff()
	fmt.Println("bf : ", rv.Elem())
	rv.Elem().Set(reflect.ValueOf(rvv))
	fmt.Println("bf : ", rv.Elem())
}

func ddd() interface{} {
	return int64(777)
}

func fff() interface{} {
	return []int{1, 2, 3}
}

func main() {
	hogehoge()
	v := []interface{}{vv}
	//v := map[interface{}]int{"a": 1, 6666: 2, "c": 3}
	var vr, vr2 []interface{}
	var sr, sr2 []interface{}

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
		err := msgpack.DeserializeAsArray(d, &sr)
		if err != nil {
			fmt.Println("des err : ", err)
		}
		fmt.Println(sr, reflect.ValueOf(sr).Type())
		d2, _ := msgpack.SerializeAsArray(sr)
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

		err := msgpack.DeserializeAsMap(d, &sr2)
		if err != nil {
			fmt.Println("des err : ", err)
		}
		fmt.Println(sr2)
		d2, _ := msgpack.SerializeAsMap(sr2)
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
