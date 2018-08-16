package main

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/shamaton/msgpack"
	a "github.com/shamaton/sandboxgo/msgpack"
	aaaa "github.com/vmihailenco/msgpack"
)

func main() {
	type st struct {
		A int
		b *uint
		c int
	}

	a.F()
	//v := []int{1, 2, 3, math.MinInt64}
	//v = nil
	//v := "this is test"
	//v := []bool{true, false}
	// v := float64(math.MaxFloat64)
	// v := []byte{0x82, 0xa1, 0x41, 0x07, 0xa1, 0x42, 0xa1, 0x37}
	//v := &st{A: math.MinInt32, b: nil}
	v := map[int]interface{}{1: 2, 3: "a", 4: []float32{1.23}}
	//v := time.Now()
	sd1, sd2 := shamaton(v)
	vd1, vd2 := vmihailenco(v)
	fmt.Println("shamaton arr : ", hex.Dump(sd1))
	fmt.Println("vmihaile arr : ", hex.Dump(vd1))
	fmt.Println("shamaton map : ", hex.Dump(sd2))
	fmt.Println("vmihaile map : ", hex.Dump(vd2))

	sss := []int{1, 3, 23, 4}
	sstest(sss)
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

	var buf bytes.Buffer
	enc := aaaa.NewEncoder(&buf).StructAsArray(true)
	err := enc.Encode(v)
	if err != nil {
		fmt.Println("err arr : ", err)
	}

	d, err := aaaa.Marshal(v)
	if err != nil {
		fmt.Println("err map : ", err)
	}

	return buf.Bytes(), d
}

func sstest(v interface{}) {
	switch v := v.(type) {
	case int:
		fmt.Println("int!! : ", v)
	case int8:
		fmt.Println("int8!! : ", v)
	case []interface{}:
		// これは無理
		fmt.Println("slice interface : ", v)
	default:
		fmt.Println("other : ", v)
	}
}

type common struct {
}

func (c *common) f() {
	fmt.Println("call common")
}

type sta struct {
	common
}

func (s *sta) f() {
	fmt.Println("call sta")
}

type stb struct {
	common
}

func (s *stb) f() {
	fmt.Println("call sta")
	s.common.f()
}
