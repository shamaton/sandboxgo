package main

import (
	"github.com/shamaton/msgpack"
	a "github.com/shamaton/sandboxgo/msgpack"
)

func main() {
	a.F()
	v := int8(-36)
	msgpack.SerializeAsArray(v)
}
