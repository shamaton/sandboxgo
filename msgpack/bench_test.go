package bench_test

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"reflect"
	"testing"

	"github.com/golang/protobuf/proto"
	shamaton "github.com/shamaton/msgpack"
	"github.com/shamaton/msgpack_bench/protocmp"
	"github.com/shamaton/zeroformatter"
	"github.com/ugorji/go/codec"
	vmihailenco "github.com/vmihailenco/msgpack"
)

type Item struct {
	ID     int
	Name   string
	Effect float32
	Num    uint
}

type User struct {
	ID       int
	Name     string
	Level    uint
	Exp      uint64
	Type     bool
	EquipIDs []uint32
	Items    []Item
}

var v = User{
	ID:       12345,
	Name:     "しゃまとん",
	Level:    99,
	Exp:      math.MaxUint32 * 2,
	Type:     true,
	EquipIDs: []uint32{1, 100, 10000, 1000000, 100000000},
	Items:    []Item{},
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

var protov = &protocmp.User{
	ID:       int32(v.ID),
	Name:     v.Name,
	Level:    uint32(v.Level),
	Exp:      v.Exp,
	Type:     v.Type,
	EquipIDs: v.EquipIDs,
	Items:    []*protocmp.Item{},
}

var _protov = &protocmp.BenchMarkStruct{
	/*
		Int:     int32(v.Int),
		Uint:    uint32(v.Uint),
		Float:   v.Float,
		Double:  v.Double,
		Bool:    v.Bool,
		String_: v.String,
		Array:   []int32{1, 2, 3, 4, 5, 6, 7, 8, 9},
		Map:     map[string]uint32{"this": 1, "is": 2, "map": 3},
		Child:   &protocmp.BenchChild{Int: 123456, String_: "this is struct of child"},
	*/
}

var (
	arrayMsgpack []byte
	mapMsgpack   []byte
	zeroFmtpack  []byte
	jsonPack     []byte
	gobPack      []byte
	protoPack    []byte
)

// for codec
var (
	mh = &codec.MsgpackHandle{}
)

func init() {
	// ugorji
	mh.MapType = reflect.TypeOf(v)

	// item
	for i := 0; i < 5; i++ {
		name := "item" + fmt.Sprint(i)
		item := Item{
			ID:     i,
			Name:   name,
			Effect: float32(i*i) / 3.0,
			Num:    uint(i * i * i * i),
		}
		v.Items = append(v.Items, item)

		pItem := &protocmp.Item{
			ID:     int32(item.ID),
			Name:   item.Name,
			Effect: item.Effect,
			Num:    uint32(item.Num),
		}
		protov.Items = append(protov.Items, pItem)
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

	d, err = zeroformatter.Serialize(v)
	if err != nil {
		fmt.Println("init err : ", err)
		os.Exit(1)
	}
	zeroFmtpack = d

	d, err = json.Marshal(v)
	if err != nil {
		fmt.Println("init err : ", err)
		os.Exit(1)
	}
	jsonPack = d

	d, err = proto.Marshal(protov)
	if err != nil {
		fmt.Println("init err : ", err)
		os.Exit(1)
	}
	protoPack = d

	buf := bytes.NewBuffer(nil)
	err = gob.NewEncoder(buf).Encode(v)
	if err != nil {
		fmt.Println("init err : ", err)
		os.Exit(1)
	}
	gobPack = buf.Bytes()
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

func BenchmarkCompareDecodeZeroformatter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var r User
		err := zeroformatter.Deserialize(&r, zeroFmtpack)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func BenchmarkCompareDecodeJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var r User
		err := json.Unmarshal(jsonPack, &r)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func BenchmarkCompareDecodeGob(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var r User
		buf := bytes.NewBuffer(gobPack)
		err := gob.NewDecoder(buf).Decode(&r)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func BenchmarkCompareDecodeProtocolBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var r protocmp.User
		err := proto.Unmarshal(protoPack, &r)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

/////////////////////////////////////////////////////////////////

func BenchmarkCompareEncodeShamaton(b *testing.B) {
	for i := 0; i < b.N; i++ {
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

		b := []byte{}
		enc := codec.NewEncoderBytes(&b, mh)
		err := enc.Encode(v)

		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func BenchmarkCompareEncodeZeroformatter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := zeroformatter.Serialize(v)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func BenchmarkCompareEncodeJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(v)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func BenchmarkCompareEncodeGob(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := bytes.NewBuffer(nil)
		err := gob.NewEncoder(buf).Encode(v)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func BenchmarkCompareEncodeProtocolBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(protov)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
