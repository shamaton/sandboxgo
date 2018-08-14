package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/shamaton/msgpack"
	a "github.com/shamaton/sandboxgo/msgpack"
	aaaa "github.com/vmihailenco/msgpack"
)

func main() {
	type st struct {
		A int
		B *uint
	}

	a.F()
	//v := []int{1, 2, 3, math.MinInt64}
	//v = nil
	//v := "this is test"
	//v := []bool{true, false}
	// v := float64(math.MaxFloat64)
	// v := []byte{0x82, 0xa1, 0x41, 0x07, 0xa1, 0x42, 0xa1, 0x37}
	// v := &st{A: math.MinInt32, B: nil}
	v := time.Now()
	d := shamaton(v)
	fmt.Println("shamaotn : ", hex.Dump(d))

	d = vmihailenco(v)
	fmt.Println("vmihailenco", hex.Dump(d))

	sss := []int{1, 3, 23, 4}
	sstest(sss)
}

func shamaton(v interface{}) []byte {
	d, err := msgpack.SerializeAsArray(v)
	if err != nil {
		fmt.Println("err : ", err)
	}
	return d
}

func vmihailenco(v interface{}) []byte {

	var buf bytes.Buffer
	enc := aaaa.NewEncoder(&buf).StructAsArray(true)
	enc.Encode(v)
	return buf.Bytes()
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
