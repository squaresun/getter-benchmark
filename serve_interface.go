package benchmark

import (
	"fmt"

	"github.com/bxcodec/faker"
	echo_interface "github.com/squaresun/getter-benchmark/echo/interface"
)

func (s SomeStruct) GetInt() int {
	return s.Int
}
func (s SomeStruct) GetInt8() int8 {
	return s.Int8
}
func (s SomeStruct) GetInt16() int16 {
	return s.Int16
}
func (s SomeStruct) GetInt32() int32 {
	return s.Int32
}
func (s SomeStruct) GetInt64() int64 {
	return s.Int64
}
func (s SomeStruct) GetString() string {
	return s.String
}
func (s SomeStruct) GetBool() bool {
	return s.Bool
}
func (s SomeStruct) GetSString() []string {
	return s.SString
}
func (s SomeStruct) GetSInt() []int {
	return s.SInt
}
func (s SomeStruct) GetSInt8() []int8 {
	return s.SInt8
}
func (s SomeStruct) GetSInt16() []int16 {
	return s.SInt16
}
func (s SomeStruct) GetSInt32() []int32 {
	return s.SInt32
}
func (s SomeStruct) GetSInt64() []int64 {
	return s.SInt64
}
func (s SomeStruct) GetSFloat32() []float32 {
	return s.SFloat32
}
func (s SomeStruct) GetSFloat64() []float64 {
	return s.SFloat64
}
func (s SomeStruct) GetSBool() []bool {
	return s.SBool
}

func serveInterface() {
	var a SomeStruct
	err := faker.FakeData(&a)
	if err != nil {
		fmt.Println(err)
	}
	echo_interface.Echo(a)
}
