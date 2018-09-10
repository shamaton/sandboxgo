package msgpack_test

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"testing"

	shamaton "github.com/shamaton/msgpack"
	"github.com/ugorji/go/codec"
	vmihailenco "github.com/vmihailenco/msgpack"
)

type Item struct {
	ID int
	/*
		Name   string
		Effect float32
		Num    uint
	*/
}

type User struct {
	/*
		ID       int
		Name     string
		Level    uint
		Exp      uint64
		Type     bool
		EquipIDs []uint32
	*/
	Items []Item
}

var v = User{
	/*
		ID:       12345,
		Name:     "しゃまとん",
		Level:    99,
		Exp:      math.MaxUint32 * 2,
		Type:     true,
		EquipIDs: []uint32{1, 100, 10000, 1000000, 100000000},
	*/
	Items: []Item{},
}

type BenchChild struct {
	Int    int
	String string
}
type BenchMarkStruct struct {
	Int    int
	Uint   uint
	Float  float32
	Double float64
	Bool   bool
	String string
	Array  []int
	Map    map[string]uint
	Child  BenchChild
}

var _v = BenchMarkStruct{
	Int:    -123,
	Uint:   456,
	Float:  1.234,
	Double: 6.789,
	Bool:   true,
	String: "this is text.",
	Array:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	Map:    map[string]uint{"this": 1, "is": 2, "map": 3},
	Child:  BenchChild{Int: 123456, String: "this is struct of child"},
}

var (
	arrayMsgpack []byte
	mapMsgpack   []byte
)

// for codec
var (
	mh = &codec.MsgpackHandle{}
)

func init() {
	// ugorji
	mh.MapType = reflect.TypeOf(v)

	// item
	for i := 0; i < 10; i++ {
		//name := "item" + fmt.Sprint(i)
		item := Item{
			ID: i,
			/*
				Name:   name,
				Effect: float32(i*i) / 3.0,
				Num:    uint(i * i * i * i),
			*/
		}
		v.Items = append(v.Items, item)
	}

	d, err := shamaton.EncodeStructAsArray(v)
	if err != nil {
		fmt.Println("init err : ", err)
		os.Exit(1)
	}
	arrayMsgpack = d
	d, err = shamaton.EncodeStructAsMap(v)
	if err != nil {
		fmt.Println("init err : ", err)
		os.Exit(1)
	}
	mapMsgpack = d

}

func TestUgorji(t *testing.T) {
	for i := 0; i < 2; i++ {
		b := []byte{}
		enc := codec.NewEncoderBytes(&b, mh)
		err := enc.Encode(v)

		if err != nil {
			t.Log(err)
		}
	}
}

func BenchmarkCompareDecodeShamaton(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var r User
		err := shamaton.DecodeStructAsMap(mapMsgpack, &r)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
func BenchmarkCompareDecodeVmihailenco(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var r User
		err := vmihailenco.Unmarshal(mapMsgpack, &r)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func BenchmarkCompareDecodeArrayShamaton(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var r User
		err := shamaton.DecodeStructAsArray(arrayMsgpack, &r)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
func BenchmarkCompareDecodeArrayVmihailenco(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var r User
		err := vmihailenco.Unmarshal(arrayMsgpack, &r)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func BenchmarkCompareDecodeUgorji(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var r User
		dec := codec.NewDecoderBytes(mapMsgpack, mh)
		err := dec.Decode(&r)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

/////////////////////////////////////////////////////////////////

func BenchmarkCompareEncodeShamaton(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := range v.Items {
			v.Items[j].ID = i
		}
		_, err := shamaton.EncodeStructAsMap(v)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func BenchmarkCompareEncodeVmihailenco(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := vmihailenco.Marshal(v)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func BenchmarkCompareEncodeArrayShamaton(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := shamaton.EncodeStructAsArray(v)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func BenchmarkCompareEncodeArrayVmihailenco(b *testing.B) {
	for i := 0; i < b.N; i++ {

		var buf bytes.Buffer
		enc := vmihailenco.NewEncoder(&buf).StructAsArray(true)
		err := enc.Encode(v)

		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func BenchmarkCompareEncodeUgorji(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := range v.Items {
			v.Items[j].ID = i
		}

		b := []byte{}
		enc := codec.NewEncoderBytes(&b, mh)
		err := enc.Encode(v)

		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
