// Package ast contains multiple functions that ease Golang Abstract Syntax Tree manipulation.
package ast

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strings"
)

type File struct {
	f *ast.File
}

// ParseFile will parse the file to be used later.
func ParseFile(filename string) (*File, error) {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, filename, nil, 0)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse file %s: %w", filename, err)
	}

	return &File{f: f}, nil
}

// AstFile will return the ast definition of file.
func (f *File) AstFile() *ast.File {
	return f.f
}

// GetStructs will return the struct definition of the specified names.
// If any of the names is not found it will return an error.
// If any of the names is not struct it will return an error.
func (f *File) GetStructs(names ...string) ([]Struct, error) {
	sdf := make([]Struct, 0, len(names))
	for _, name := range names {
		spec := f.FindTypeSpecByName(name)
		if spec == nil {
			return nil, fmt.Errorf("there is not definition of %s", name)
		}

		df, err := extractStruct(spec)
		if err != nil {
			return nil, err
		}
		sdf = append(sdf, df)
	}

	return sdf, nil
}

// GetImports will find every import declaration in the file.
func (f *File) GetImports() []Import {
	nf := &NodeFinder{Selector: FindImportSpec()}
	ast.Walk(nf, f.f)

	var res []Import
	for _, node := range nf.Results {
		is, ok := node.(*ast.ImportSpec)
		if !ok { // impossible but let it be for now.
			continue
		}

		var name string
		if is.Name != nil {
			name = is.Name.Name
		}

		res = append(res, Import{
			Name: name,
			Path: is.Path.Value,
		})
	}

	return res
}

// FindTypeSpecByName will return *ast.TypeSpec of the type definition that have the specific name.
func (f *File) FindTypeSpecByName(name string) *ast.TypeSpec {
	nf := &NodeFinder{Selector: FindTypeSpecByName(name)}
	ast.Walk(nf, f.f)

	if ts, ok := nf.Result().(*ast.TypeSpec); ok {
		return ts
	} else {
		return nil
	}
}

func extractStruct(node *ast.TypeSpec) (Struct, error) {
	var res Struct

	structType, ok := node.Type.(*ast.StructType)
	if !ok {
		return res, fmt.Errorf("%v is not struct", node.Name)
	}

	var fields []StructField
	for _, f := range structType.Fields.List {
		field, err := extractField(f)
		if err != nil {
			return res, err
		}

		if field != nil {
			fields = append(fields, *field)
		}
	}

	return Struct{
		Name:   node.Name.Name,
		Fields: fields,
	}, nil
}

func extractField(node *ast.Field) (*StructField, error) {
	// Anonymous type is skipped
	if len(node.Names) == 0 {
		return nil, nil
	}

	name := node.Names[0].Name
	if !ast.IsExported(name) {
		return nil, nil
	}

	typeName, err := getTypeName(node.Type)
	if err != nil {
		return nil, fmt.Errorf("failed to extract field %s: %w", name, err)
	}

	return &StructField{
		Name:      name,
		Type:      typeName,
		ZeroValue: getZeroValue(typeName),
	}, nil
}

func getZeroValue(tipe string) string {
	// List in https://go.dev/tour/basics/11
	switch tipe {
	case "float32", "float64",
		"int", "int8", "int16", "int32", "int64", "uint",
		"uint8", "uint16", "uint32", "uint64", "uintptr",
		"complex64", "complex128",
		"byte", "rune":
		return "0"
	case "string":
		return `""`
	case "bool":
		return "false"
	default:
		if strings.HasPrefix(tipe, "map") || strings.HasPrefix(tipe, "[") || strings.HasPrefix(tipe, "*") {
			return "nil"
		}

		return fmt.Sprintf("%s{}", tipe)
	}
}

func getTypeName(expr ast.Expr) (string, error) {
	switch tipe := expr.(type) {
	case *ast.Ident:
		return tipe.Name, nil
	case *ast.ArrayType:
		if tipe.Len != nil {
			return "", errors.New("array field type is not supported yet")
		}

		tName, err := getTypeName(tipe.Elt)
		return fmt.Sprintf("[]%s", tName), err
	case *ast.StarExpr:
		tName, err := getTypeName(tipe.X)
		return fmt.Sprintf("*%s", tName), err
	case *ast.SelectorExpr:
		tName, err := getTypeName(tipe.X)
		return fmt.Sprintf("%s.%s", tName, tipe.Sel.Name), err
	case *ast.MapType:
		kName, err := getTypeName(tipe.Key)
		if err != nil {
			return "", err
		}

		vName, err := getTypeName(tipe.Value)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("map[%s]%s", kName, vName), nil
	default:
		return "", fmt.Errorf("unrecognized expression type: %s", reflect.TypeOf(expr))
	}
}
