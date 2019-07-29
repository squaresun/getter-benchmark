package benchmark

import (
	"fmt"

	"github.com/bxcodec/faker"
	echo_struct "github.com/squaresun/getter-benchmark/echo/struct"
)

type SomeStruct struct {
	Int      int
	Int8     int8
	Int16    int16
	Int32    int32
	Int64    int64
	String   string
	Bool     bool
	SString  []string
	SInt     []int
	SInt8    []int8
	SInt16   []int16
	SInt32   []int32
	SInt64   []int64
	SFloat32 []float32
	SFloat64 []float64
	SBool    []bool
}

func serveStruct() {
	var a SomeStruct
	err := faker.FakeData(&a)
	if err != nil {
		fmt.Println(err)
	}
	echo_struct.Echo(echo_struct.SomeStruct{
		Int:      a.Int,
		Int8:     a.Int8,
		Int16:    a.Int16,
		Int32:    a.Int32,
		Int64:    a.Int64,
		String:   a.String,
		Bool:     a.Bool,
		SString:  a.SString,
		SInt:     a.SInt,
		SInt8:    a.SInt8,
		SInt16:   a.SInt16,
		SInt32:   a.SInt32,
		SInt64:   a.SInt64,
		SFloat32: a.SFloat32,
		SFloat64: a.SFloat64,
		SBool:    a.SBool,
	})
}
