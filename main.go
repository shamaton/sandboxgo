package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math"

	"github.com/shamaton/msgpack"
	a "github.com/shamaton/sandboxgo/msgpack"
	aaaa "github.com/vmihailenco/msgpack"
)

func main() {
	a.F()
	v := []int{1, 2, 3, math.MinInt64}
	d, err := msgpack.SerializeAsArray(v)
	fmt.Println(err)
	fmt.Println(hex.Dump(d))

	var buf bytes.Buffer
	enc := aaaa.NewEncoder(&buf).StructAsArray(true)
	enc.Encode(math.MaxUint32)

	var result int32
	aaaa.Unmarshal(buf.Bytes(), &result)
	fmt.Println(">>>>>>>>>>", math.MaxUint32, result)

}
