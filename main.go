package main

import (
	"github.com/shamaton/msgpack"
	a "github.com/shamaton/sandboxgo/msgpack"
)

func main() {
	a.F()
	v := int16(-1)
	msgpack.SerializeAsArray(v)
}
