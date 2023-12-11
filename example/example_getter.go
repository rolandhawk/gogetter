// Code generated by gogetter. DO NOT EDIT.
// Source: example.go

// Package example is generated by gogetter.
package example

import (
	"os"

	"golang.org/x/tools/imports"
)

// Getter for BasicType

func (x *BasicType) GetBool() (o bool) {
	if x != nil {
		return x.Bool
	}

	return o
}

func (x *BasicType) GetString() (o string) {
	if x != nil {
		return x.String
	}

	return o
}

func (x *BasicType) GetByte() (o byte) {
	if x != nil {
		return x.Byte
	}

	return o
}

func (x *BasicType) GetRune() (o rune) {
	if x != nil {
		return x.Rune
	}

	return o
}

func (x *BasicType) GetInt() (o int) {
	if x != nil {
		return x.Int
	}

	return o
}

func (x *BasicType) GetInt8() (o int8) {
	if x != nil {
		return x.Int8
	}

	return o
}

func (x *BasicType) GetInt16() (o int16) {
	if x != nil {
		return x.Int16
	}

	return o
}

func (x *BasicType) GetInt32() (o int32) {
	if x != nil {
		return x.Int32
	}

	return o
}

func (x *BasicType) GetInt64() (o int64) {
	if x != nil {
		return x.Int64
	}

	return o
}

func (x *BasicType) GetUint() (o uint) {
	if x != nil {
		return x.Uint
	}

	return o
}

func (x *BasicType) GetUint8() (o uint8) {
	if x != nil {
		return x.Uint8
	}

	return o
}

func (x *BasicType) GetUint16() (o uint16) {
	if x != nil {
		return x.Uint16
	}

	return o
}

func (x *BasicType) GetUint32() (o uint32) {
	if x != nil {
		return x.Uint32
	}

	return o
}

func (x *BasicType) GetUint64() (o uint64) {
	if x != nil {
		return x.Uint64
	}

	return o
}

func (x *BasicType) GetUintptr() (o uintptr) {
	if x != nil {
		return x.Uintptr
	}

	return o
}

func (x *BasicType) GetFloat32() (o float32) {
	if x != nil {
		return x.Float32
	}

	return o
}

func (x *BasicType) GetFloat64() (o float64) {
	if x != nil {
		return x.Float64
	}

	return o
}

func (x *BasicType) GetComplex64() (o complex64) {
	if x != nil {
		return x.Complex64
	}

	return o
}

func (x *BasicType) GetComplex128() (o complex128) {
	if x != nil {
		return x.Complex128
	}

	return o
}

// Getter for Complex

func (x *Complex) GetStruct() (o BasicType) {
	if x != nil {
		return x.Struct
	}

	return o
}

func (x *Complex) GetPointer() (o *BasicType) {
	if x != nil {
		return x.Pointer
	}

	return o
}

func (x *Complex) GetStdLibStruct() (o os.File) {
	if x != nil {
		return x.StdLibStruct
	}

	return o
}

func (x *Complex) GetStdLibPointer() (o *os.File) {
	if x != nil {
		return x.StdLibPointer
	}

	return o
}

func (x *Complex) GetForeignStruct() (o imports.Options) {
	if x != nil {
		return x.ForeignStruct
	}

	return o
}

func (x *Complex) GetForeignPointer() (o *imports.Options) {
	if x != nil {
		return x.ForeignPointer
	}

	return o
}

func (x *Complex) GetMap() (o map[BasicType]string) {
	if x != nil {
		return x.Map
	}

	return o
}

func (x *Complex) GetSlice() (o []*BasicType) {
	if x != nil {
		return x.Slice
	}

	return o
}

func (x *Complex) GetAlias() (o Alias) {
	if x != nil {
		return x.Alias
	}

	return o
}

func (x *Complex) GetUnderlying() (o Underlying) {
	if x != nil {
		return x.Underlying
	}

	return o
}
