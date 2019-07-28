package echo_interface

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

func Echo(s SomeInterface) SomeInterface {
	a := SomeStruct{
		Int:      s.GetInt(),
		Int8:     s.GetInt8(),
		Int16:    s.GetInt16(),
		Int32:    s.GetInt32(),
		Int64:    s.GetInt64(),
		String:   s.GetString(),
		Bool:     s.GetBool(),
		SString:  s.GetSString(),
		SInt:     s.GetSInt(),
		SInt8:    s.GetSInt8(),
		SInt16:   s.GetSInt16(),
		SInt32:   s.GetSInt32(),
		SInt64:   s.GetSInt64(),
		SFloat32: s.GetSFloat32(),
		SFloat64: s.GetSFloat64(),
		SBool:    s.GetSBool(),
	}
	return a
}
