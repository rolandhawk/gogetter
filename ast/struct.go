package ast

// Struct contain information of the struct that we are interested.
type Struct struct {
	Name   string        // Struct name
	Fields []StructField // Field definition
}

type StructField struct {
	Name      string // Name of the field
	Type      string // Type of the field
	ZeroValue string // Zero value for the type
}

type Import struct {
	Name string // Import name, may be nil
	Path string // Path of the import
}
