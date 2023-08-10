package ast

import (
	"go/ast"
)

// NodeSelector will return true if the node is the expected node.
type NodeSelector func(ast.Node) bool

// FindTypeSpecByName is a selector for typespec ast.Node.
func FindTypeSpecByName(name string) NodeSelector {
	return func(n ast.Node) bool {
		typespec, ok := n.(*ast.TypeSpec)
		if !ok {
			return false
		}

		return typespec.Name.Name == name
	}
}

// FindImportSpec is a selector for typespec ast.Node.
func FindImportSpec() NodeSelector {
	return func(n ast.Node) bool {
		_, ok := n.(*ast.ImportSpec)
		return ok
	}
}

// NodeFinder implement go/ast.Visitor
type NodeFinder struct {
	// The desired ast.Node will be stored here
	Results []ast.Node

	// Selector for Visit function
	Selector NodeSelector
}

// Visit will invoke NodeFinder.Selector. If it's return true,
// it will store the node into the NodeFinder.Res.
func (nf *NodeFinder) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return nf
	}

	if nf.Selector(node) {
		nf.Results = append(nf.Results, node)
		return nil
	}

	return nf
}

// Result is syntatic sugar to get the first element of nf.Results.
// It will return nil if nf.Results is empty.
func (nf *NodeFinder) Result() ast.Node {
	if len(nf.Results) > 0 {
		return nf.Results[0]
	}
	return nil
}
