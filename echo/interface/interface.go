package echo_interface

type SomeInterface interface {
	GetInt() int
	GetInt8() int8
	GetInt16() int16
	GetInt32() int32
	GetInt64() int64
	GetString() string
	GetBool() bool
	GetSString() []string
	GetSInt() []int
	GetSInt8() []int8
	GetSInt16() []int16
	GetSInt32() []int32
	GetSInt64() []int64
	GetSFloat32() []float32
	GetSFloat64() []float64
	GetSBool() []bool
}
