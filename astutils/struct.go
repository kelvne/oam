package astutils

import "go/ast"

// ExtractResources returns structs on a *ast.File
func ExtractResources(f *ast.File) map[string]*Struct {
	m := make(map[string]*Struct)
	for _, decl := range f.Decls {
		if typeDecl, ok := decl.(*ast.GenDecl); ok {
			if s := parseStructSpec(typeDecl); s != nil {
				m[s.Name] = s
			}
		}
	}
	return m
}
