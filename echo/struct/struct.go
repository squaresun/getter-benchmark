package echo_struct

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

func Echo(s SomeStruct) SomeStruct {
	a := SomeStruct{
		Int:      s.Int,
		Int8:     s.Int8,
		Int16:    s.Int16,
		Int32:    s.Int32,
		Int64:    s.Int64,
		String:   s.String,
		Bool:     s.Bool,
		SString:  s.SString,
		SInt:     s.SInt,
		SInt8:    s.SInt8,
		SInt16:   s.SInt16,
		SInt32:   s.SInt32,
		SInt64:   s.SInt64,
		SFloat32: s.SFloat32,
		SFloat64: s.SFloat64,
		SBool:    s.SBool,
	}
	return a
}
