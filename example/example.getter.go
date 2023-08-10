// Code generated by gogetter. DO NOT EDIT.
// Source: example.go

// Package example is generated by gogetter.
package example

import (
	"os"

	"golang.org/x/tools/imports"
)

// Getter for BasicType

func (x *BasicType) GetBool() bool {
	if x != nil {
		return x.Bool
	}

	return false
}

func (x *BasicType) GetString() string {
	if x != nil {
		return x.String
	}

	return ""
}

func (x *BasicType) GetByte() byte {
	if x != nil {
		return x.Byte
	}

	return 0
}

func (x *BasicType) GetRune() rune {
	if x != nil {
		return x.Rune
	}

	return 0
}

func (x *BasicType) GetInt() int {
	if x != nil {
		return x.Int
	}

	return 0
}

func (x *BasicType) GetInt8() int8 {
	if x != nil {
		return x.Int8
	}

	return 0
}

func (x *BasicType) GetInt16() int16 {
	if x != nil {
		return x.Int16
	}

	return 0
}

func (x *BasicType) GetInt32() int32 {
	if x != nil {
		return x.Int32
	}

	return 0
}

func (x *BasicType) GetInt64() int64 {
	if x != nil {
		return x.Int64
	}

	return 0
}

func (x *BasicType) GetUint() uint {
	if x != nil {
		return x.Uint
	}

	return 0
}

func (x *BasicType) GetUint8() uint8 {
	if x != nil {
		return x.Uint8
	}

	return 0
}

func (x *BasicType) GetUint16() uint16 {
	if x != nil {
		return x.Uint16
	}

	return 0
}

func (x *BasicType) GetUint32() uint32 {
	if x != nil {
		return x.Uint32
	}

	return 0
}

func (x *BasicType) GetUint64() uint64 {
	if x != nil {
		return x.Uint64
	}

	return 0
}

func (x *BasicType) GetUintptr() uintptr {
	if x != nil {
		return x.Uintptr
	}

	return 0
}

func (x *BasicType) GetFloat32() float32 {
	if x != nil {
		return x.Float32
	}

	return 0
}

func (x *BasicType) GetFloat64() float64 {
	if x != nil {
		return x.Float64
	}

	return 0
}

func (x *BasicType) GetComplex64() complex64 {
	if x != nil {
		return x.Complex64
	}

	return 0
}

func (x *BasicType) GetComplex128() complex128 {
	if x != nil {
		return x.Complex128
	}

	return 0
}

// Getter for Complex

func (x *Complex) GetStruct() BasicType {
	if x != nil {
		return x.Struct
	}

	return BasicType{}
}

func (x *Complex) GetPointer() *BasicType {
	if x != nil {
		return x.Pointer
	}

	return nil
}

func (x *Complex) GetStdLibStruct() os.File {
	if x != nil {
		return x.StdLibStruct
	}

	return os.File{}
}

func (x *Complex) GetStdLibPointer() *os.File {
	if x != nil {
		return x.StdLibPointer
	}

	return nil
}

func (x *Complex) GetForeignStruct() imports.Options {
	if x != nil {
		return x.ForeignStruct
	}

	return imports.Options{}
}

func (x *Complex) GetForeignPointer() *imports.Options {
	if x != nil {
		return x.ForeignPointer
	}

	return nil
}

func (x *Complex) GetMap() map[BasicType]string {
	if x != nil {
		return x.Map
	}

	return nil
}

func (x *Complex) GetSlice() []*BasicType {
	if x != nil {
		return x.Slice
	}

	return nil
}

func (x *Complex) GetAlias() Alias {
	if x != nil {
		return x.Alias
	}

	return DefaultAlias
}

func (x *Complex) GetUnderlying() Underlying {
	if x != nil {
		return x.Underlying
	}

	return 0
}