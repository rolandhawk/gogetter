//go:generate gogetter -out test.go BasicType Complex

package example

import (
	"os"

	"golang.org/x/tools/imports"
)

type Alias = int32
type Underlying int32

type BasicType struct {
	Bool   bool
	String string
	Byte   byte
	Rune   rune

	Int   int
	Int8  int8
	Int16 int16
	Int32 int32
	Int64 int64

	Uint    uint
	Uint8   uint8
	Uint16  uint16
	Uint32  uint32
	Uint64  uint64
	Uintptr uintptr

	Float32 float32
	Float64 float64

	Complex64  complex64
	Complex128 complex128
}

type Complex struct {
	Struct  BasicType
	Pointer *BasicType

	StdLibStruct  os.File
	StdLibPointer *os.File

	ForeignStruct  imports.Options
	ForeignPointer *imports.Options

	Map   map[BasicType]string
	Slice []*BasicType

	Alias      Alias
	Underlying Underlying
}
