package msgpack

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math"

	"github.com/vmihailenco/msgpack"
)

// F is func
func F() {
	p("int 1 -------------->")
	msg(0)
	p("int 2 -------------->")
	msg(-32)
	p("int 3 -------------->")
	msg(-36)
	p("int min -------------->")
	msg(int64(math.MinInt32))
	p("array -------------->")
	msg([]int{})
	p("array 2 -------------->")
	msg([]int{1, 2, 3, math.MinInt64})
	p("map -------------->")
	msg(map[int]int{1: 2, 3: 4, 5: math.MaxInt32})
	p("string -------------->")
	msg("this is test")
	p("string 2-------------->")
	msg("01234567890123456789012345678901")
	p("string empty-------------->")
	msg("")

	p("nil -------------->")
	msg(nil)

	p("st 1 -------------->")
	type st1 struct {
		A int
	}
	msg(st1{A: 7})
	p("st 2-------------->")
	type st2 struct {
		A int
		B string
	}
	msg(st2{A: 7, B: "7"})

	p("st 3-------------->")
	var nilSt2 *st2
	nilSt2 = nil
	msg(nilSt2)
	p("-------------->")
	// p("-------------->")

}

func msg(v ...interface{}) {
	b, e := msgpack.Marshal(v...)
	printDump("def", b, e)

	var buf bytes.Buffer
	enc := msgpack.NewEncoder(&buf).StructAsArray(true)
	err := enc.Encode(v...)
	printDump("arr", buf.Bytes(), err)
}

func printDump(t string, b []byte, e error) {
	if e != nil {
		fmt.Println("error : ", e)
		return
	}
	fmt.Println(t, ":")
	fmt.Println(hex.Dump(b))
	fmt.Println("")
}

func p(v ...interface{}) {
	fmt.Println(v...)
}
